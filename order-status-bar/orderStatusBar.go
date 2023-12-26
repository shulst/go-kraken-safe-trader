package order_status_bar

import (
	progress_bar "github.com/shulst/go-kraken-safe-trader/progress-bar"
	"image"
	"image/color"
	"image/draw"
)

type Colors interface {
	BG() color.RGBA
	FG() color.RGBA
	Overlay() color.RGBA
	FGOverlay() color.RGBA
}

type OrderType string

const (
	Buy  = OrderType("buy")
	Sell = OrderType("sell")
)

type OrderStatusBar struct {
	ProgressBar progress_bar.ProgressBar
	OrderType   OrderType
}

func Create(progressBar progress_bar.ProgressBar, orderType OrderType) OrderStatusBar {
	return OrderStatusBar{
		ProgressBar: progressBar,
		OrderType:   orderType,
	}
}
func drawRectangle(img draw.Image, color color.Color, x1, y1, x2, y2 int) {
	for i := x1; i < x2; i++ {
		img.Set(i, y1, color)
		img.Set(i, y2-1, color)
	}

	for i := y1; i <= y2; i++ {
		img.Set(x1, i, color)
		img.Set(x2-1, i, color)
	}
}

func (bar OrderStatusBar) Draw(width int, height int) image.Image {
	black := color.RGBA{A: 255}
	grey := color.RGBA{R: 110, G: 110, B: 110, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	green := color.RGBA{G: 164, A: 255}
	red := color.RGBA{R: 164, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw white color on whole image
	draw.Draw(img, img.Bounds(), &image.Uniform{C: white}, image.Point{}, draw.Src)

	var fg color.RGBA
	if bar.OrderType == Buy {
		fg = green
	} else {
		fg = red
	}

	// Add FG rectangle
	var fgRect image.Rectangle
	offset := (width / int(bar.ProgressBar.Size)) * int(bar.ProgressBar.Fills[1])
	if bar.OrderType == Buy {
		fgRect = image.Rect(width-offset, 2, width-2, height-1)
	} else {
		fgRect = image.Rect(1, 1, offset, height-1)
	}
	draw.Draw(img, fgRect, &image.Uniform{C: fg}, image.Point{}, draw.Src)
	drawRectangle(img, white, fgRect.Min.X, fgRect.Min.X, fgRect.Max.X, fgRect.Max.Y)

	// Offset
	var overlayRect image.Rectangle
	overlayOffset := (width / int(bar.ProgressBar.Size)) * int(bar.ProgressBar.Fills[0])
	if bar.ProgressBar.Fills[0] > 0 {
		if bar.OrderType == Buy {
			overlayRect = image.Rect(width/2, 4, width/2+overlayOffset/2, height-4)
		} else {
			overlayRect = image.Rect(width/2-overlayOffset/2, 4, width/2, height-4)
		}
		draw.Draw(img, overlayRect, &image.Uniform{C: grey}, image.Point{}, draw.Src)
		drawRectangle(img, black, overlayRect.Min.X, overlayRect.Min.X, overlayRect.Max.X, overlayRect.Max.Y)
	}

	drawRectangle(img, fg, 0, 0, width, height)
	return img
}
