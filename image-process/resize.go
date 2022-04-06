package imgProc

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

func ResizeImage(fileName string, b []byte, size string) {
	width := strings.Split(strings.ToLower(size), "x")[0]
	w, _ := strconv.Atoi(width)
	height := strings.Split(strings.ToLower(size), "x")[1]
	h, _ := strconv.Atoi(height)

	// Later on, this will be used whenever
	// periodically deleting images folder
	// will be added
	if _, err := os.Stat("./images/"); os.IsNotExist(err) {
		err = os.Mkdir("images", 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	output, err := os.Create("./images/" + fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer output.Close()

	src, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	switch strings.Split(fileName, ".")[1] {
	case "jpeg", "jpg":
		jpeg.Encode(output, dst, nil)
	case "png":
		png.Encode(output, dst)
	}

}
