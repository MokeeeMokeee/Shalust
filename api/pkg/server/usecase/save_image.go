package usecase

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func SaveIconImage(data string, user_id string) (string, error) {

	encodeData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.Write(encodeData)
	tmpFile.Close()

	file, err := os.Open(tmpFile.Name())
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return "", err
	}

	if format == "jpeg" {
		result_path, err := Upload_S3(tmpFile.Name(), fmt.Sprintf("user_icons/%s.jpg", user_id))
		return result_path, err
	} else if format == "png" {
		result_path, err := Upload_S3(tmpFile.Name(), fmt.Sprintf("user_icons/%s.png", user_id))
		return result_path, err
	}

	return "", nil
}
