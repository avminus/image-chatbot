package controllers

import (
	"encoding/json"
	"fmt"
	"image-chatbot/models"
	structs "image-chatbot/structures"
	"image-chatbot/utils"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func SendTextMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var promptMessage structs.TextMessage

	err := json.NewDecoder(r.Body).Decode(&promptMessage)
	message := structs.Message{Role: "user", Type: "text", Data: promptMessage}
	if err != nil {
		http.Error(w, "error while parsing request body! ", http.StatusBadRequest)
		return
	}

	res, err := models.SendTextMessage(&message)

	if err != nil {
		http.Error(w, fmt.Sprintf("error while fetching response from AI assistant! %v", err), http.StatusBadRequest)
		return
	}

	responseText := structs.TextMessage{Content: res}
	responseMessage := structs.Message{Role: "assistant", Type: "text", Id: uuid.New().String(), Data: responseText}

	var messages []structs.TextMessage
	messages = append(messages, responseText)

	messageResponse := structs.MessageResponse{Status: true,
		Data: messages,
	}

	err = json.NewEncoder(w).Encode(messageResponse)

	if err != nil {
		http.Error(w, "error while parsing response body ", http.StatusInternalServerError)
		return
	}

	structs.MessageList = append(structs.MessageList,
		message,
		responseMessage,
	)

}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	if !utils.ImageSupported(contentType) {
		http.Error(w, "Media type not supported! ", http.StatusUnsupportedMediaType)
		return
	}

	message := structs.Message{Role: "user", Type: contentType}
	utils.ParseImageMessage(r, &message)

	err := models.UploadImage(&message)

	if err != nil {
		http.Error(w, "Error in uploading the file ", http.StatusInternalServerError)
		return
	}

	structs.ImageIdToTypeMap[message.Id] = contentType
	structs.MessageList = append(structs.MessageList, message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image Uploaded Successfully!"))
}

func GetImageById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	imageId := params["id"]

	imageMessage, err := models.GetImageById(imageId)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error while fetching Image! %v", err), http.StatusInternalServerError)
		fmt.Printf("%v", err)
		return
	}

	w.Header().Set("Content-Type", structs.ImageIdToTypeMap[imageId])
	_, err = w.Write(imageMessage.Content)

	if err != nil {
		http.Error(w, "Error sending Image", http.StatusInternalServerError)
		fmt.Printf("%v", err)
		return
	}

}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	pNo, err := strconv.ParseInt(r.URL.Query().Get("pNo"), 0, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while parsing the Page Number: %v", err), http.StatusBadRequest)
		return
	}

	pSize, err := (strconv.ParseInt(r.URL.Query().Get("pSize"), 0, 0))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while parsing the Page Size: %v", err), http.StatusBadRequest)
		return
	}

	messageList, err := models.GetAllMessages(int(pNo), int(pSize))
	messageResponse := structs.MetaMessageResponse{Status: true, Data: *messageList}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error while fetching the messages: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(messageResponse)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error while encoding the messages: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
