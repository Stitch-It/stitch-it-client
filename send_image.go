package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func sendProcessedImageToServer(fileName string, client Client) {
	b, w := createMultiPartFormData("image", fileName)

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.7:3030/images", &b)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := client.http.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(string(bytes))
}

func createMultiPartFormData(fieldName, fileName string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error

	w := multipart.NewWriter(&b)
	var fw io.Writer
	file := openFile("./images/" + fileName)

	fw, err = w.CreateFormFile(fieldName, file.Name())
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	w.Close()
	return b, w
}

func openFile(fileName string) *os.File {
	r, _ := os.Open(fileName)
	return r
}
