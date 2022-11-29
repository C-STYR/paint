package ui

import (
	"paint/apptype"
	"paint/paintcanvas"
	"paint/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PaintCanvas *paintcanvas.PaintCanvas
	PaintWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
