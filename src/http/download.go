package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getFileNameFromHeader(header http.Header) (string, error) {
	headerKey := http.CanonicalHeaderKey("content-disposition")
	if len(header[headerKey]) == 0 {
		return "", fmt.Errorf("Could not find the content-disposition header, are you downloading a file?")
	}
	headerValue := header[headerKey][0]
	for _, v := range strings.Split(headerValue, ";") {
		keyvalue := strings.TrimSpace(v)
		if strings.HasPrefix(keyvalue, "filename") {
			return strings.Trim(strings.Split(keyvalue, "=")[1], "\""), nil
		}
	}
	return "", fmt.Errorf("Could not find a suggested filename")
}

func Download(url string, savePath string) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Could not fetch the file: %s\n", err)
	}
	defer resp.Body.Close()
	outFileName, err := getFileNameFromHeader(resp.Header)
	if err != nil {
		return err
	}
	outFile, err := os.Create(filepath.Join(savePath, outFileName))
	if err != nil {
		return fmt.Errorf("Could not create file to write: %s\n", err)
	}

	_, err = io.Copy(outFile, resp.Body)
	return err
}
