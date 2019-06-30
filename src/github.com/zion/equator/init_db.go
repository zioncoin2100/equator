package equator

import (
	"github.com/zion/go/support/db"
	"github.com/zion/equator/db2/core"
	"github.com/zion/equator/db2/history"
	"github.com/zion/equator/log"
)

func initEquatorDb(app *App) {
	session, err := db.Open("postgres", app.config.DatabaseURL)

	if err != nil {
		log.Panic(err)
	}
	session.DB.SetMaxIdleConns(4)
	session.DB.SetMaxOpenConns(12)

	app.historyQ = &history.Q{session}
}

func initCoreDb(app *App) {
	session, err := db.Open("postgres", app.config.ZionCoreDatabaseURL)

	if err != nil {
		log.Panic(err)
	}

	session.DB.SetMaxIdleConns(4)
	session.DB.SetMaxOpenConns(12)
	app.coreQ = &core.Q{session}
}

func init() {
	appInit.Add("equator-db", initEquatorDb, "app-context", "log")
	appInit.Add("core-db", initCoreDb, "app-context", "log")
}
