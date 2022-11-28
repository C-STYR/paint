package main

import (
	"image/color"
	"paint/apptype"
	"paint/swatch"
	"paint/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	paintApp := app.New()

	paintWindow := paintApp.NewWindow("paint")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PaintWindow: paintWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PaintWindow.ShowAndRun()
}
