package utils

import (
	"fmt"
	structs "image-chatbot/structures"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const IMAGE_DIRECTORY = "/resources/images/"

func ImageSupported(format string) bool {
	switch format {
	case "image/jpeg":
		return true
	case "image/png":
		return true
	case "image/gif":
		return true
	default:
		return false
	}
}

func GetImagePath(imageId string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Error in getting current directory: %v", err)
	}

	return filepath.Dir(wd) + IMAGE_DIRECTORY + imageId, nil
}

func ParseImageMessage(r *http.Request, msg *structs.Message) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	imageMessage := structs.ImageMessage{Content: body}
	msg.Data = imageMessage
	return nil
}
