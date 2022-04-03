package main

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

func resizeImage(fileName string, b []byte, size string) []byte {
	width := strings.Split(strings.ToLower(size), "x")[0]
	w, _ := strconv.Atoi(width)
	height := strings.Split(strings.ToLower(size), "x")[1]
	h, _ := strconv.Atoi(height)

	var img []byte

	switch strings.Split(fileName, ".")[1] {
	case "jpeg", "jpg":
		img = resizeJpeg(fileName, b, w, h)
	case "png":
		img = resizePng(fileName, b, w, h)
	}

	return img
}

func resizeJpeg(fileName string, b []byte, w int, h int) []byte {
	if _, err := os.Stat("./images/"); os.IsNotExist(err) {
		err = os.Mkdir("images", 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	err := os.Chdir("images")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	src, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	jpeg.Encode(output, dst, nil)

	data, _ := os.ReadFile(output.Name())

	err = os.Chdir("..")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	println("image processed successfully")

	return data
}

func resizePng(fileName string, b []byte, w int, h int) []byte {
	if _, err := os.Stat("./images/"); os.IsNotExist(err) {
		err = os.Mkdir("images", 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	err := os.Chdir("images")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	src, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	png.Encode(output, dst)

	data, _ := os.ReadFile(output.Name())

	err = os.Chdir("..")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	println("image processed successfully")

	return data
}
