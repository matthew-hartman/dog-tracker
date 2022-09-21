package main

import (
	"encoding/json"
	"fmt"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type Vapid struct {
	Subscription webpush.Subscription `json:"subscription"`
}

type Notification struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Vibrate []int  `json:"vibrate"`
	Icon    string `json:"icon"`
	DogID   string `json:"dogID"`
}

func (c *Controller) handleNotify(ctx echo.Context) error {
	log.Println(ctx.Request().URL)
	DogID := ctx.QueryParam("dogID")

	collection, err := c.App.Dao().FindCollectionByNameOrId("dogs")
	if err != nil {
		return ctx.JSON(500, err)
	}

	records, err := c.App.Dao().FindRecordsByExpr(collection, dbx.HashExp{"id": DogID})
	if err != nil {
		return ctx.JSON(500, err)
	}

	if len(records) == 0 {
		return ctx.JSON(404, "Dog not found")
	}

	data := records[0].Data()
	if data["state"] != "notify" {
		return ctx.JSON(400, "Bad request")
	}

	records[0].SetDataValue("state", "acked")

	err = c.App.Dao().SaveRecord(records[0])
	if err != nil {
		return ctx.JSON(500, err)
	}

	return ctx.JSON(200, "OK")
}

func (c *Controller) handleNotifyState(e *core.RecordUpdateEvent) error {
	data := e.Record.Data()
	if data["state"] != "notify" {
		return nil
	}

	collection, err := c.App.Dao().FindCollectionByNameOrId("vapid")
	if err != nil {
		return err
	}

	records, err := c.App.Dao().FindRecordsByExpr(collection, dbx.HashExp{"user": data["user"]})
	if err != nil {
		return err
	}

	for _, record := range records {

		raw, err := record.MarshalJSON()
		if err != nil {
			return err
		}

		vapid := Vapid{}
		err = json.Unmarshal(raw, &vapid)
		if err != nil {
			return err
		}

		notification := Notification{
			Title:   "Zoomies",
			Body:    fmt.Sprintf("%s's owner is here!", data["name"]),
			Vibrate: []int{100, 50, 300},
			Icon:    "/icon.png",
			DogID:   e.Record.Id,
		}
		notificationJSON, err := json.Marshal(notification)
		if err != nil {
			return err
		}

		resp, err := webpush.SendNotification(
			notificationJSON,
			&vapid.Subscription,
			&webpush.Options{
				VAPIDPublicKey:  vapidPublicKey,
				VAPIDPrivateKey: vapidPrivateKey,
				TTL:             30,
			})
		if err != nil {
			log.Println(err)
			return err
		}
		defer resp.Body.Close()
		log.Println(resp.StatusCode)

		if resp.StatusCode == 410 {
			c.App.Dao().DeleteRecord(record)
		}
	}

	return nil
}
