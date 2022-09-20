package main

import (
	"encoding/json"
	"fmt"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
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
		}
		notificationJSON, err := json.Marshal(notification)
		if err != nil {
			return err
		}

		resp, err := webpush.SendNotification(
			notificationJSON,
			&vapid.Subscription,
			&webpush.Options{
				//Subscriber:      "test@example.com",
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
