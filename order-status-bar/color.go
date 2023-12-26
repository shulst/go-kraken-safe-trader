package order_status_bar

import (
	"image/color"
)

type Color struct {
	R, G, B, A float32
}

func RGBAToColor(c color.RGBA) Color {
	return Color{
		R: float32(c.R) / 255,
		G: float32(c.G) / 255,
		B: float32(c.B) / 255,
		A: float32(c.A) / 255,
	}
}

func (c Color) RGBA() color.RGBA {
	return color.RGBA{
		R: uint8(c.R * 255),
		G: uint8(c.G * 255),
		B: uint8(c.B * 255),
		A: uint8(c.A * 255),
	}
}

func CombineColors(colors ...color.RGBA) color.RGBA {
	bg := RGBAToColor(colors[0])
	for _, fgRGBA := range colors[1:] {
		fg := RGBAToColor(fgRGBA)
		var r Color
		r.A = 1 - (1-fg.A)*(1-bg.A)
		r.R = fg.R*fg.A/r.A + bg.R*bg.A*(1-fg.A)/r.A
		r.G = fg.G*fg.A/r.A + bg.G*bg.A*(1-fg.A)/r.A
		r.B = fg.B*fg.A/r.A + bg.B*bg.A*(1-fg.A)/r.A

		bg = r
	}
	return bg.RGBA()
}
