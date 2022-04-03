package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func sendProcessedImageToServer(b []byte, fileName string, client Client) {
	buf := bytes.NewBuffer(b)

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.7:3030/images", buf)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

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
