package main

import (
	"fmt"
	"os"
)

func processImage(fileName string) error {
	file, err := os.Open(fmt.Sprintf("images/%s", fileName))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	defer file.Close()

	// img, _, _ := image.Decode(file)

	// resizing will go here

	// colArr, colNum := resizeImage(img, size)

	// if (len(colArr) != 0) && (colNum != nil) {
	// 	if errMsg := generateExcelPattern(imgNm, colArr, colNum); errMsg != "" {
	// 		genErrMsg := errMsg

	// 		genErr = c.JSON(fiber.Map{"status": 500, "message": genErrMsg})
	// 	}
	// } else {
	// 	genErr = c.JSON(fiber.Map{"status": 201, "message": "Image successfully processed and excel pattern generated"})
	// }

	return nil
}
