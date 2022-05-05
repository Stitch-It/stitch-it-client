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

	// check to see if the patterns directory is there for saving patterns to
	// and create it if it isn't
	if _, err := os.Stat("./patterns/"); os.IsNotExist(err) {
		err = os.Mkdir("patterns", 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	// grab the resized image
	imgFile, err := os.Open("./images/" + fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// decode it into an image.Image (can think of it as a wrapper around the bitmap (2D byte array) of the raster
	// image that provides functionality to retrieve the RGBA value from each byte)
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// close the image as we no longer need it open
	err = imgFile.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// change to the patterns directory to save the pattern to
	err = os.Chdir("./patterns")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// retrieve the filename minus the extension so we can later change it, but have a consistent filename
	fileNameXlsxExtension := strings.Split(fileName, ".")[0]

	// represent the width and height of the image.Image cleaner
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	// create a new excel workbook
	patternFile := excelize.NewFile()

	defer patternFile.Close()

	// set the first workbook sheet's name to List (will hold the list of numbers and the corresponding thread color names)
	patternFile.SetSheetName("Sheet1", "List")
	// create a new workbook sheet and change its name to Pattern (will obviously hold the generated pattern)
	patternSheet := patternFile.NewSheet("Pattern")
	println(patternSheet)

	// this hasn't been completed yet. For some reason, it isn't setting the style. But wasn't important, so did not complete yet.
	cellStyle, err := patternFile.NewStyle(`"border":[{"type":"left","color":"000000","style":2},{"type":"top","color":"000000","style":2},{"type":"bottom","color":"000000","style":2},{"type":"right","color":"000000","style":2}]`)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// generate the pattern and grab the list of numbers and corresponding thread color names to create the color List
	colorMap := generatePatternSheet(img, patternFile, cellStyle, width, height)

	// generate the Color List from the colorMap
	generateColorListSheet(colorMap, patternFile, cellStyle)

	// save the workbook
	err = patternFile.SaveAs(fileNameXlsxExtension + "@" + authorScreenName + ".xlsx")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func generatePatternSheet(img image.Image, patternFile *excelize.File, cellStyle int, width, height int) map[string]int {
	// cross stitch patterns rarely have 0s, so make sure the first color is set to be represented by 1
	colorNumber := 1

	// initialize a colorMap to hold the list of numbers and corresponding thread color names
	colorMap := make(map[string]int)

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

			// initalize a bank of thread colors to test each pixel's color against
			// using the syke99/go-c2dmc package
			colorBank := dmc.NewColorBank()

			// calculate the closest matching thread color to the pixel's RGB values
			color, _ := colorBank.RgbToDmc(r, g, b)

			// create a cell in the excel sheet (since the bounds of an image.Image start at
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
				colorMap[color] = 1
			} else {
				if _, ok := colorMap[color]; !ok {
					colorNumber++
					colorMap[color] = colorNumber
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
func generateColorListSheet(colorMap map[string]int, patternFile *excelize.File, cellStyle int) {
	for clr, nmb := range colorMap {
		patternFile.SetCellValue("List", "A"+strconv.Itoa(nmb), nmb)
		patternFile.SetCellStyle("List", "A"+strconv.Itoa(nmb), "A"+strconv.Itoa(nmb), cellStyle)
		patternFile.SetCellValue("List", "B"+strconv.Itoa(nmb), clr)
		patternFile.SetCellStyle("List", "B"+strconv.Itoa(nmb), "B"+strconv.Itoa(nmb), cellStyle)
	}
}
