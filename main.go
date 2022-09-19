package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"os"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

//go:generate npm run build

//go:embed dist
var f embed.FS

var (
	vapidPublicKey  = os.Getenv("VAPID_PUB")
	vapidPrivateKey = os.Getenv("VAPID_PRI")
)

func main() {
	f, err := fs.Sub(f, "dist")
	if err != nil {
		log.Fatal(err)
	}

	app := pocketbase.New()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(f, false))
		return nil
	})
	app.OnRecordAfterUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		data := e.Record.Data()
		if data["state"] != "notify" {
			return nil
		}
		collection, err := app.Dao().FindCollectionByNameOrId("vapid")
		if err != nil {
			return err
		}

		record, err := app.Dao().FindFirstRecordByData(collection, "user", data["user"])
		if err != nil {
			return err
		}

		type Vapid struct {
			Subscription webpush.Subscription `json:"subscription"`
		}

		raw, err := record.MarshalJSON()
		if err != nil {
			return err
		}

		vapid := Vapid{}
		err = json.Unmarshal(raw, &vapid)
		if err != nil {
			return err
		}

		resp, err := webpush.SendNotification([]byte("Test"), &vapid.Subscription, &webpush.Options{
			Subscriber:      "test@example.com",
			VAPIDPublicKey:  vapidPublicKey,
			VAPIDPrivateKey: vapidPrivateKey,
			TTL:             30,
		})
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
