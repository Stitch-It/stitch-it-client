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

func GenerateExcelPattern(fileName, authorScreenName string) {

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

	fileNameXlsxExtension := strings.Split(fileName, ".")[0]

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	patternFile := excelize.NewFile()

	defer patternFile.Close()

	patternFile.SetSheetName("Sheet1", "List")
	patternSheet := patternFile.NewSheet("Pattern")
	println(patternSheet)

	cellStyle, err := patternFile.NewStyle(`"border":[{"type":"left","color":"000000","style":2},{"type":"top","color":"000000","style":2},{"type":"bottom","color":"000000","style":2},{"type":"right","color":"000000","style":2}]`)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	colorMap := generatePatternSheet(img, patternFile, cellStyle, width, height)

	generateColorListSheet(colorMap, patternFile, cellStyle)

	err = patternFile.SaveAs(fileNameXlsxExtension + "@" + authorScreenName + ".xlsx")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func generatePatternSheet(img image.Image, patternFile *excelize.File, cellStyle int, width, height int) map[string]int {
	colorNumber := 1

	colorMap := make(map[string]int)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			rgbaColor := img.At(x, y)

			r_, g_, b_, _ := rgbaColor.RGBA()

			r := float64(r_ / 255)
			g := float64(g_ / 255)
			b := float64(b_ / 255)

			colorBank := dmc.NewColorBank()

			color, _ := colorBank.RgbToDmc(r, g, b)

			cellName, err := excelize.CoordinatesToCellName(x+1, y+1)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			if cellName == "A1" {
				colorMap[color] = 1
			} else {
				if _, ok := colorMap[color]; !ok {
					colorNumber++
					colorMap[color] = colorNumber
				}
			}

			patternFile.SetCellValue("Pattern", cellName, colorMap[color])
		}
	}

	return colorMap
}

func generateColorListSheet(colorMap map[string]int, patternFile *excelize.File, cellStyle int) {
	for clr, nmb := range colorMap {
		patternFile.SetCellValue("List", "A"+strconv.Itoa(nmb), nmb)
		patternFile.SetCellStyle("List", "A"+strconv.Itoa(nmb), "A"+strconv.Itoa(nmb), cellStyle)
		patternFile.SetCellValue("List", "B"+strconv.Itoa(nmb), clr)
		patternFile.SetCellStyle("List", "B"+strconv.Itoa(nmb), "B"+strconv.Itoa(nmb), cellStyle)
	}
}
