package main

// This is where the heavy lifting of generating the
// pattern will take place
func generatePattern(fileName string, size string) {
	// width := strings.Split(strings.ToLower(size), "x")[0]
	// w, _ := strconv.Atoi(width)
	// height := strings.Split(strings.ToLower(size), "x")[1]
	// h, _ := strconv.Atoi(height)

	// file, err := os.Open(fmt.Sprintf("./images/%s", fileName))
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// defer file.Close()

	// procImg, _, err := image.DecodeConfig(file)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// img, _, err := image.Decode(file)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// if (w == procImg.Width) && (h == procImg.Height) {
	// 	// Generate pdf
	// }

}

// func generatePdf(pdfWidth int, width int, height int, image image.Image) {
// 	m := pdf.NewMaroto(consts.Portrait, consts.A4)
// 	m.SetPageMargins(20, 10, 20)
// 	buildPatternTable(pdfWidth, width, height, image)
// }

// func buildPatternTable(pdfWidth int, width int, height int, image image.Image) {
// 	buildTableContent2DArray(width, height, image)
// }

// func buildTableContent2DArray(width int, height int, image image.Image) [][]int {
// colorMap := make(map[string]int)
// for {

// }
// return nil
// }
