package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
	colorA, colorB uint8
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w , i.h)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	v := uint8((x+y)/2)
	return color.RGBA{R: v, G: v, B: i.colorA, A: i.colorB}
}

func main() {
	m := &Image{100, 75, 128, 177}

	pic.ShowImage(m)
}