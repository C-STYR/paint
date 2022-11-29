package main

import (
	"image/color"
	"paint/apptype"
	"paint/paintcanvas"
	"paint/swatch"
	"paint/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	paintApp := app.New()

	paintWindow := paintApp.NewWindow("paint")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	paintCanvasConfig := apptype.PaintCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	paintCanvas := paintcanvas.NewPxCanvas(&state, paintCanvasConfig)

	appInit := ui.AppInit{
		PaintCanvas: paintCanvas,
		PaintWindow: paintWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PaintWindow.ShowAndRun()
}
