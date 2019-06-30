package equator

import (
	"github.com/zion/equator/reap"
)

func initReaper(app *App) {
	app.reaper = reap.New(app.config.HistoryRetentionCount, app.EquatorSession(nil))
}

func init() {
	appInit.Add("reaper", initReaper, "app-context", "log", "equator-db")
}
