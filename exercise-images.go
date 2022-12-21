package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	colorModel color.Model
	bounds     image.Rectangle
}

func (image Image) ColorModel() color.Model {
	return image.colorModel
}

func (image Image) Bounds() image.Rectangle {
	return image.bounds
}

func (image Image) At(dx, dy int) color.Color {
	value := uint8((dx + dy)/2)
	return color.RGBA{value, value, 255, 255}
}

func main() {
	m := Image{
		color.RGBAModel,
		image.Rect(0, 0, 100, 100),
	}
	pic.ShowImage(m)
}
