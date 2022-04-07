package gen

import (
	"fmt"
	"image"
	"os"

	dmc "github.com/syke99/go-c2dmc"
	"github.com/tealeg/xlsx/v2"
)

func GenerateExcelPattern(fileName, authorScreenName string) string {
	if _, err := os.Stat("./patterns/"); os.IsNotExist(err) {
		err = os.Mkdir("patterns", 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	imgFile, err := os.Open("./images/" + fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	path := "./patterns/" + fileName

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	println(path)

	wb := xlsx.NewFile()

	patternSheet, err := wb.AddSheet("Pattern")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	cellStyle := xlsx.NewStyle()
	cellStyle.Alignment.Horizontal = "center"
	cellStyle.Alignment.Vertical = "center"
	cellStyle.Border.Left = "solid"
	cellStyle.Border.LeftColor = "black"
	cellStyle.Border.Right = "solid"
	cellStyle.Border.RightColor = "black"
	cellStyle.Border.Top = "solid"
	cellStyle.Border.TopColor = "black"
	cellStyle.Border.Bottom = "solid"
	cellStyle.Border.BottomColor = "black"
	cellStyle.Font.Bold = true
	cellStyle.ApplyAlignment = true
	cellStyle.ApplyBorder = true
	cellStyle.ApplyFont = true

	colorMap := generatePatternSheet(img, patternSheet, cellStyle, width, height)

	colorListSheet, err := wb.AddSheet("Color List")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	generateColorListSheet(colorMap, cellStyle, colorListSheet)

	err = wb.Save(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return path
}

func generatePatternSheet(image image.Image, patternSheet *xlsx.Sheet, cellStyle *xlsx.Style, width, height int) map[string]int {
	colorNumber := 1

	colorMap := make(map[string]int)

	for x := 0; x < width; x++ {
		row := patternSheet.AddRow()
		for y := 0; y < height; y++ {
			cell := row.AddCell()

			rgbaColor := image.At(width, height)

			colorBank := dmc.NewColorBank()

			r, g, b := colorBank.RgbA(rgbaColor)

			color, _ := colorBank.Rgb(r, g, b)

			if x == 0 && y == 0 {
				colorMap[color] = colorNumber
				cell.SetInt(colorNumber)
				cell.SetStyle(cellStyle)
				colorNumber++
			} else {
				if _, ok := colorMap[color]; !ok {
					colorMap[color] = colorNumber
				}
				cell.SetInt(colorNumber)
				cell.SetStyle(cellStyle)
				colorNumber++
			}
		}
	}

	return colorMap
}

func generateColorListSheet(colorMap map[string]int, cellStyle *xlsx.Style, colorListSheet *xlsx.Sheet) {
	for clr, nmb := range colorMap {
		row := colorListSheet.AddRow()

		numberCell := row.AddCell()
		numberCell.SetInt(nmb)
		numberCell.SetStyle(cellStyle)

		colorCell := row.AddCell()
		colorCell.SetString(clr)
		colorCell.SetStyle(cellStyle)
	}
}
