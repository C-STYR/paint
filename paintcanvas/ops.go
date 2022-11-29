package paintcanvas

import "fyne.io/fyne/v2"

func (paintCanvas *PaintCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	paintCanvas.CanvasOffset.X += xDiff
	paintCanvas.CanvasOffset.Y += yDiff
	paintCanvas.Refresh()
}

func (paintCanvas *PaintCanvas) Scale(direction int) {
	switch {
	case direction > 0:
		paintCanvas.PxSize += 1
	case direction < 0:
		if paintCanvas.PxSize > 2 {
			paintCanvas.PxSize -= 1
		}
	default:
		paintCanvas.PxSize = 10
	}
}