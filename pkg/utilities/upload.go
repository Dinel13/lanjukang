package utilities

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadFile uploads a file to the server
func UploadedImage(uploadedImage multipart.File, header *multipart.FileHeader) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// change woking direktory in server
	// becouse in server the cwd is in /
	if dir != "/home/din/project/Lanjukang-be" {
		dir = "/var/www/wisata"
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(header.Filename))
	fileLocation := filepath.Join(dir, "images", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedImage); err != nil {
		return "", err
	}

	return filename, nil

}

// DeleteImage is a function to delete image
func DeleteImage(filename string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// change woking direktory in server
	// becouse in server the cwd is in /
	if dir != "/home/din/project/Lanjukang-be" {
		dir = "/var/www/wisata"
	}

	fileLocation := filepath.Join(dir, "images", filename)
	err = os.Remove(fileLocation)
	if err != nil {
		return err
	}

	return nil
}
