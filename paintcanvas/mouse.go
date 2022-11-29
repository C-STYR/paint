package paintcanvas

import (
	"paint/paintcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func(paintCanvas *PaintCanvas) Scrolled(ev *fyne.ScrollEvent) {
	paintCanvas.Scale(int(ev.Scrolled.DY))
	paintCanvas.Refresh()
}

func (paintCanvas *PaintCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := paintCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(paintCanvas.appState, paintCanvas, ev)
	}
	paintCanvas.TryPan(paintCanvas.mouseState.previousCoord, ev)
	paintCanvas.Refresh()
	paintCanvas.mouseState.previousCoord = &ev.PointEvent
}


func (paintCanvas *PaintCanvas) MouseIn(ev *desktop.MouseEvent) {}

func (paintCanvas *PaintCanvas) MouseOut() {}

func (paintCanvas *PaintCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(paintCanvas.appState, paintCanvas, ev)
}

func (paintCanvas *PaintCanvas) MouseUp(ev *desktop.MouseEvent) {}