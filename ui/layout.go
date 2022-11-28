package ui

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.PaintWindow.SetContent(swatchesContainer)
}