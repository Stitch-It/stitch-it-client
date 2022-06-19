package gen

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	dmc "github.com/syke99/go-c2dmc"
	"github.com/xuri/excelize/v2"
	"image"
	"strconv"
)

func GenerateExcelPattern(img *image.RGBA) string {
	var buf bytes.Buffer

	// represent the width and height of the image.Image cleaner
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	// create a new Excel workbook
	patternFile := excelize.NewFile()

	defer patternFile.Close()

	// set the first workbook sheet's name to List (will hold the list of numbers and the corresponding thread color names)
	patternFile.SetSheetName("Sheet1", "List")
	// create a new workbook sheet and change its name to Pattern (will obviously hold the generated pattern)
	patternSheet := patternFile.NewSheet("Pattern")
	println(patternSheet)

	// generate the pattern and grab the list of numbers and corresponding thread color names to create the color List
	colorMap := generatePatternSheet(img, patternFile, width, height)

	// generate the Color List from the colorMap
	generateColorListSheet(colorMap, patternFile)

	patternFile.Write(&buf)

	encodedFileString := b64.StdEncoding.EncodeToString(buf.Bytes())

	return encodedFileString
}

type threadInfo struct {
	colorNumber int
	colorFloss  string
}

func generatePatternSheet(img image.Image, patternFile *excelize.File, width, height int) map[string]threadInfo {
	// cross stitch patterns rarely have 0s, so make sure the first color is set to be represented by 1
	colorNumber := 1

	// initialize a colorMap to hold the list of numbers and corresponding thread color names
	colorMap := make(map[string]threadInfo)

	// initialize a bank of thread colors to test each pixel's color against
	// using the syke99/go-c2dmc package
	colorBank := dmc.NewColorBank()
	
	// loop through the image.Image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			// grab the color.Color from each pixel (byte)
			rgbaColor := img.At(x, y)

			// convert that color.Color to its RGBA representation
			r_, g_, b_, _ := rgbaColor.RGBA()

			// remove the alpha value (A) from the RGBA
			// to return it to an RGB value
			r := float64(r_ / 255)
			g := float64(g_ / 255)
			b := float64(b_ / 255)

			// calculate the closest matching thread color to the pixel's RGB values
			color, floss := colorBank.RgbToDmc(r, g, b)

			// create a cell in the Excel sheet (since the bounds of an image.Image start at
			// 0 and 0 (width and height, respectfully), we need to increment each pixel's number
			// by one so that excelize knows where to create a cell (otherwise, you cut off one row
			// and one column of pixels)
			cellName, err := excelize.CoordinatesToCellName(x+1, y+1)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			// generate the colorMap based on if it's the first color to be found,
			// or whether the color already exists in the color map
			if cellName == "A1" {
				initialColor := threadInfo{
					colorNumber: 1,
					colorFloss:  floss,
				}
				colorMap[color] = initialColor
			} else {
				if _, ok := colorMap[color]; !ok {
					nextColor := threadInfo{
						colorNumber: colorNumber,
						colorFloss:  floss,
					}
					colorNumber++
					colorMap[color] = nextColor
				}
			}

			// place the nuber corresponding to the nearest matching thread color to
			// the cell it belongs to whenever mapped to an excel sheet
			patternFile.SetCellValue("Pattern", cellName, colorMap[color])
		}
	}

	return colorMap
}

// loops through the colorMap and adds a new entry for each number/corresponding thread color
func generateColorListSheet(colorMap map[string]threadInfo, patternFile *excelize.File) {
	for color, info := range colorMap {
		patternFile.SetCellValue("List", "A"+strconv.Itoa(info.colorNumber), info.colorNumber)
		patternFile.SetCellValue("List", "B"+strconv.Itoa(info.colorNumber), color)
		patternFile.SetCellValue("List", "C"+strconv.Itoa(info.colorNumber), info.colorFloss)
	}
}
