package gen

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	dmc "github.com/syke99/go-c2dmc"
	"github.com/xuri/excelize/v2"
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

	err = imgFile.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = os.Chdir("./patterns")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fileNameNoExtension := strings.Split(fileName, ".")[0]

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	patternFile := excelize.NewFile()

	patternFile.DeleteSheet("Sheet1")

	patternSheet := patternFile.NewSheet("Pattern")
	println(patternSheet)
	colorListSheet := patternFile.NewSheet("List")
	println(colorListSheet)

	colorMap := generatePatternSheet(img, patternFile, width, height)

	generateColorListSheet(colorMap, patternFile)

	err = patternFile.SaveAs(fileNameNoExtension + ".xlsx")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return "./patterns/" + fileNameNoExtension
}

func generatePatternSheet(image image.Image, patternFile *excelize.File, width, height int) map[string]int {
	colorNumber := 1

	colorMap := make(map[string]int)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			rgbaColor := image.At(x, y)

			rgbaColor.RGBA()

			colorBank := dmc.NewColorBank()

			r, g, b := colorBank.RgbA(rgbaColor)

			l_, a_, b_ := colorBank.RgbToLab(r, g, b)

			color, _ := colorBank.LabToDmc(l_, a_, b_)

			cellName, err := excelize.CoordinatesToCellName(x+1, y+1)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			if cellName == "A1" {
				colorMap[color] = 1
				patternFile.SetCellValue("Pattern", cellName, colorNumber)
				colorNumber++
			} else {
				if _, ok := colorMap[color]; !ok {
					colorMap[color] = colorNumber
					patternFile.SetCellValue("Pattern", cellName, colorNumber)
					colorNumber++
				} else {
					patternFile.SetCellValue("Pattern", cellName, colorNumber)
				}
			}
		}
	}

	return colorMap
}

func generateColorListSheet(colorMap map[string]int, patternFile *excelize.File) {
	for clr, nmb := range colorMap {
		patternFile.SetCellValue("List", "A"+strconv.Itoa(nmb), nmb)
		patternFile.SetCellValue("List", "B"+strconv.Itoa(nmb), clr)
	}
}
