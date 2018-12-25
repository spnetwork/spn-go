package horizon

func initSpnCoreInfo(app *App) {
	app.UpdateSpnCoreInfo()
	return
}

func init() {
	appInit.Add("spnCoreInfo", initSpnCoreInfo, "app-context", "log")
}
