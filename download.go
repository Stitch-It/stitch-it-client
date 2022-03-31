package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func downloadImage(URL string, user string) (string, error) {
	fileName := createFileName(URL, user)

	res, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New("received non 200 response code")
	}

	if _, err := os.Stat("./images/"); os.IsNotExist(err) {
		err = os.Mkdir("images", 0755)
		if err != nil {
			return "", err
		}
	}

	err = os.Chdir("images")
	if err != nil {
		return "", err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}

	err = os.Chdir("..")
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func createFileName(URL string, user string) string {
	uniqueId := uuid.New()

	uu := strings.Replace(uniqueId.String(), "-", "", -1)

	uniqueUser := uu + user

	splits := strings.Split(URL, ".")

	fileExt := "." + splits[len(splits)-1]

	return uniqueUser + fileExt
}
