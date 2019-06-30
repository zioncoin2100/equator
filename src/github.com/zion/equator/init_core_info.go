package equator

func initZionCoreInfo(app *App) {
	app.UpdateZionCoreInfo()
	return
}

func init() {
	appInit.Add("zionCoreInfo", initZionCoreInfo, "app-context", "log")
}
