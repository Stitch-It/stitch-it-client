package imgProc

import (
	"bytes"
	"fmt"
	"golang.org/x/image/draw"
	"image"
)

func ResizeImage(b interface{}, metric bool, width int, height int) *image.RGBA {
	var w int
	var h int

	// determine how many pixels to set dimensions to based on
	// value of metric. There are 6 stitches per centimeter, or
	// 14 stitches per inch
	switch metric {
	case false:
		w = width * 14
		h = height * 14
	case true:
		w = width * 6
		h = height * 6
	}

	// decode the bytes of the image
	src, _, err := image.Decode(bytes.NewReader(b.([]byte)))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// this will create a new *image.RGBA for drawing the resize image to
	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	// resize the image by using the CatmullRom resizing algorithm. It is the slowest
	// resizing algorithm, but produces the highest quality results. dst is the *image.RGBA to draw do,
	// dst.Rect is the dimensions we want it to be, src is the originally sized image, src.Bounds is the
	// originally sized image's dimensions, draw.Over and opts nil aren't important (draw.Over basically
	// means don't mask any part of the dst's output)
	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	return dst
}
