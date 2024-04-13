package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

func SaveImage(encode string, img_type string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll("images", os.ModePerm)
	if err != nil {
		return "", err
	}

	filename := filepath.Join("images", fmt.Sprintf("%s_%d.png", img_type, os.Getpid()))

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(decoded)
	if err != nil {
		return "", err
	}

	return filename, nil
}
