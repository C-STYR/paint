package paintcanvas

import (
	"image"
	"image/color"
	"paint/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type PaintCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PaintCanvas struct {
	widget.BaseWidget
	apptype.PaintCanvasConfig
	renderer    *PaintCanvasRenderer
	PixelData   image.Image
	mouseState  PaintCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (paintCanvas *PaintCanvas) Bounds() image.Rectangle {
	x0 := int(paintCanvas.CanvasOffset.X)
	y0 := int(paintCanvas.CanvasOffset.Y)
	x1 := int(paintCanvas.PxCols*paintCanvas.PxSize + int(paintCanvas.CanvasOffset.X))
	y1 := int(paintCanvas.PxRows*paintCanvas.PxSize + int(paintCanvas.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
		pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		pos.Y < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PaintCanvasConfig) *PaintCanvas {
	paintCanvas := &PaintCanvas{
		PaintCanvasConfig: config,
		appState:          state,
	}
	paintCanvas.PixelData = NewBlankImage(paintCanvas.PxCols, paintCanvas.PxRows, color.NRGBA{128, 128, 128, 255})
	paintCanvas.ExtendBaseWidget(paintCanvas)

	return paintCanvas
}

func (paintCanvas *PaintCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(paintCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PaintCanvasRenderer{
		paintCanvas:  paintCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}
	paintCanvas.renderer = renderer
	return renderer
}

func (paintCanvas *PaintCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		paintCanvas.Pan(*previousCoord, ev.PointEvent)
	}
}


// Brushable interface
func (paintCanvas *PaintCanvas) SetColor(c color.Color, x, y int) {
	if nrgba, ok := paintCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}
	if rgba, ok := paintCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}
	paintCanvas.Refresh()
}

func (paintCanvas *PaintCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := paintCanvas.Bounds()
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}

	pxSize := float32(paintCanvas.PxSize)
	xOffset := paintCanvas.CanvasOffset.X
	yOffset := paintCanvas.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}