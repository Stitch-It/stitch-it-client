package imgProc

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func DownloadImage(URL string) (string, []byte, error) {
	// this just creates a unique filename for each image
	fileName := createFileName(URL)

	// download the image
	res, err := http.Get(URL)
	if err != nil {
		return "", nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", nil, errors.New("received non 200 response code")
	}

	// read in the bytes of the image
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	// return those bytes
	return fileName, bytes, nil
}

func createFileName(URL string) string {
	u := uuid.New()

	uniqueId := strings.ReplaceAll(u.String(), "-", "")

	splits := strings.Split(URL, ".")

	fileExt := "." + splits[len(splits)-1]

	return uniqueId + fileExt
}
