package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(
		// Top: eventually top side,
		nil,
		// Bottom:
		swatchesContainer,
		// Left:
		nil,
		// Right: Color Picker
		colorPicker,
		// Center Object
		app.PixlCanvas,
	)

	app.PixlWindow.SetContent(appLayout)
}
