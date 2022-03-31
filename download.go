package main

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func downloadImage(URL string) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	err = os.Chdir("images")
	if err != nil {
		return err
	}

	file, err := os.Create("testfile.jpeg")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}
