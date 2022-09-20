package main

import (
	"embed"
	"io/fs"
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

//go:generate npm run build

//go:embed dist
var f embed.FS

var (
	vapidPublicKey  = os.Getenv("VAPID_PUBLIC_KEY")
	vapidPrivateKey = os.Getenv("VAPID_PRIVATE_KEY")
)

type Controller struct {
	App *pocketbase.PocketBase
}

func main() {
	f, err := fs.Sub(f, "dist")
	if err != nil {
		log.Fatal(err)
	}

	app := pocketbase.New()
	ctlr := &Controller{App: app}

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(f, false))
		return nil
	})
	app.OnRecordAfterUpdateRequest().Add(ctlr.handleNotifyState)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
