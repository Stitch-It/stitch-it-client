package imgProc

import (
	"bytes"
	"fmt"
	"golang.org/x/image/draw"
	"image"
)

//func ResizeImage(fileName string, b []byte, size string) {
//	// size was extracted from the tweet's "data section" but will be turned into textbox inputs (restricted to whole digit inputs) along with a checkbox for either cm. or in.
//	width := strings.Split(strings.ToLower(size), "x")[0]
//	w, _ := strconv.Atoi(width)
//	height := strings.Split(strings.ToLower(size), "x")[1]
//	h, _ := strconv.Atoi(height)
//
//	curDir, err := os.Getwd()
//	if err != nil {
//		fmt.Printf("err: %v\n", err)
//	}
//
//	// because of the nature of each tweet being ran in its own thread, sometimes, you'll be in the patterns directory
//	// because the pattern hasn't finished generating, so to save the image in the correct directory, we check to see if
//	// we are in the patterns directory and then cd out of it if so
//	if strings.HasSuffix(curDir, "patterns") {
//		err = os.Chdir("..")
//		if err != nil {
//			fmt.Printf("err: %v\n", err)
//		}
//	}
//
//	// Later on, this will be used whenever
//	// periodically deleting images folder
//	// will be added
//	if _, err := os.Stat("./images/"); os.IsNotExist(err) {
//		err = os.Mkdir("images", 0755)
//		if err != nil {
//			fmt.Printf("err: %v\n", err)
//		}
//	}
//
//	// create a file to save the resized image to
//	output, err := os.Create("./images/" + fileName)
//	if err != nil {
//		fmt.Printf("err: %v\n", err)
//	}
//	defer output.Close()
//
//	// decode the bytes of the image
//	src, _, err := image.Decode(bytes.NewReader(b))
//	if err != nil {
//		fmt.Printf("err: %v\n", err)
//	}
//
//	// this will create a new *image.RGBA for drawing the resize image to
//	dst := image.NewRGBA(image.Rect(0, 0, w, h))
//
//	// resize the image by using the CatmullRom resizing algorithm. It is the slowest
//	// resizing algorithm, but produces the highest quality results. dst is the *image.RGBA to draw do,
//	// dst.Rect is the dimensions we want it to be, src is the originally sized image, src.Bounds is the
//	// originally sized image's dimensions, draw.Over and opts nil aren't important (draw.Over basically
//	// means don't mask any part of the dst's output)
//	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
//
//	// determine how to encode the resized image based on the original image's filetype (.jpeg/.jpg or .png)
//	switch strings.Split(fileName, ".")[1] {
//	case "jpeg", "jpg":
//		jpeg.Encode(output, dst, nil)
//	case "png":
//		png.Encode(output, dst)
//	}
//}

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
