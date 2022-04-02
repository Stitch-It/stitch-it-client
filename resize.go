package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

func resizeImage(fileName string, size string) error {
	file, err := os.Open(fmt.Sprintf("./images/%s", fileName))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	width := strings.Split(strings.ToLower(size), "x")[0]
	w, _ := strconv.Atoi(width)
	height := strings.Split(strings.ToLower(size), "x")[1]
	h, _ := strconv.Atoi(height)

	defer file.Close()

	switch strings.Split(fileName, ".")[1] {
	case "jpeg", "jpg":
		{
			err := os.Chdir("images")
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			output, err := os.Create("processed_" + fileName)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			src, _ := jpeg.Decode(file)

			dst := image.NewRGBA(image.Rect(0, 0, w, h))

			draw.ApproxBiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

			jpeg.Encode(output, dst, nil)

			err = os.Chdir("..")
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			println("image processed successfully")
		}
	case "png":
		{
			err := os.Chdir("images")
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			output, err := os.Create("processed_" + fileName)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			src, _ := png.Decode(file)

			dst := image.NewRGBA(image.Rect(0, 0, w, h))

			draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

			png.Encode(output, dst)

			err = os.Chdir("..")
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			println("image processed successfully")
		}
	}

	return nil
}