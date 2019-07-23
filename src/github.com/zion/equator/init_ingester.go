package equator

import (
	"log"

	"github.com/zion/equator/ingest"
)

func initIngester(app *App) {
	if !app.config.Ingest {
		return
	}

	if app.networkPassphrase == "" {
		log.Fatal("Cannot start ingestion without network passphrase.  Please confirm connectivity with zion-core.")
	}

	app.ingester = ingest.New(
		app.networkPassphrase,
		app.config.ZioncoreURL,
		app.CoreSession(nil),
		app.EquatorSession(nil),
	)

	app.ingester.SkipCursorUpdate = app.config.SkipCursorUpdate
}

func init() {
	appInit.Add("ingester", initIngester, "app-context", "log", "equator-db", "core-db", "ZioncoreInfo")
}
