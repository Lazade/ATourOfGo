package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{
	x, y, w, h int
}

func (img Image) Bounds() image.Rectangle {
	x := img.x
	y := img.y
	w := img.w
	h := img.h
	return image.Rect(x, y, w, h)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	generator := func(i, j int) uint8 {
		return uint8(i^j)
	}
	v := generator(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{0, 0, 100, 100}
	pic.ShowImage(m)
}
