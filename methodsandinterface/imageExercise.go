package methodsandinterface

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

//Image is a struct
type Image struct {
	Width, Height int
	colr          uint8
}

//Bounds needs a comment
func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

//ColorModel needs a comment
func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

//At needs a comment
func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colr + uint8(x), r.colr + uint8(y), 255, 255}
}

//ImageExersize needs a comment
func ImageExersize() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
