package models

import (
	"bytes"
	"fmt"
	"image-chatbot/client"
	structs "image-chatbot/structures"
	"image-chatbot/utils"
	"io"
	"os"

	"github.com/google/uuid"
)

const (
	IMAGE_DIRECTORY = "/resources/images/"
)

func GetImageById(imageId string) (*structs.ImageMessage, error) {

	imagePath, err := utils.GetPath(imageId, IMAGE_DIRECTORY)
	if err != nil {
		return nil, fmt.Errorf("error in getting the file path: %v", err)
	}

	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("error in reading the file: %v", err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("error reading the file into the buffer: %v", err)
	}

	return &structs.ImageMessage{Content: buffer}, nil
}

func UploadImage(image *structs.Message) error {
	imageID := uuid.New().String()
	image.Id = imageID

	filepath, err := utils.GetPath(imageID, IMAGE_DIRECTORY)

	if err != nil {
		return fmt.Errorf("error in getting the file path: %v", err)
	}

	outPutFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error while creating directory|path: %v", err)
	}
	defer outPutFile.Close()

	switch msg := image.Data.(type) {
	case structs.ImageMessage:
		_, err = io.Copy(outPutFile, bytes.NewReader(msg.Content))
		if err != nil {
			return fmt.Errorf("error while saving image: %v", err)
		}
	default:
		return fmt.Errorf("unknown message type")
	}

	if err != nil {
		return fmt.Errorf("error while saving image: %v", err)
	}

	return nil
}

func GetAllMessages(pageNo, pageSize int) (*structs.Messages, error) {
	if pageNo <= 0 || (pageNo-1)*pageSize > len(structs.MessageList) {
		return nil, fmt.Errorf("pagenumber and pagesize combination invalid, messageListSize: %v", len(structs.MessageList))
	}

	lastIndex := (len(structs.MessageList) - 1 - ((pageNo - 1) * pageSize))
	count := 0
	var messages structs.Messages
	for i := lastIndex; count < pageSize && i >= 0; i-- {
		message := structs.MessageList[i]
		message.Data = struct{}{}
		messages.MessageList = append(messages.MessageList, message)
		count++
	}

	return &messages, nil
}

func SendTextMessage(text *structs.Message) (string, error) {
	textID := uuid.New().String()
	text.Id = textID
	var promptText string

	switch msg := text.Data.(type) {
	case structs.TextMessage:
		promptText = msg.Content
	default:
		return "", fmt.Errorf("unknown message type")
	}

	response, err := client.GenerateText(promptText)

	if err != nil {
		return "", fmt.Errorf("Erorr in fetching response from LLM api : %v \n", err)
	}

	return response, nil
}
