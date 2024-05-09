package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// Saves the image to the filesystem and returns the path
// img_type refers to the purpose of the image (avatar / post)
func SaveImage(encode string, img_type string) (string, error) {
	// Extract the base64 part without the "data:image/png;base64," prefix
	if len(encode) == 0 {
		return "", errors.New("no image")
	}
	base64Data := strings.Split(encode, ";base64,")[1]
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	syscall.Umask(0)
	err = os.MkdirAll("images", 0o777)
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

// EncodeImageToBase64 reads an image file and returns its base64 encoded content
func EncodeImage(imagePath string) (string, error) {
	// Read image file
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
	}

	// Encode to base64 string
	base64Str := base64.StdEncoding.EncodeToString(imageBytes)

	InfoConsoleLog(base64Str)

	return base64Str, nil
}
