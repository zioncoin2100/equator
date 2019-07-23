package equator

func initZioncoreInfo(app *App) {
	app.UpdateZioncoreInfo()
	return
}

func init() {
	appInit.Add("ZioncoreInfo", initZioncoreInfo, "app-context", "log")
}
