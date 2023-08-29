package routes

import (
	"image-chatbot/controllers"

	"github.com/gorilla/mux"
)

func RegisterChatBotRoutes(router *mux.Router) {
	router.HandleFunc("/message", controllers.SendTextMessage).Methods("POST")
	router.HandleFunc("/image", controllers.UploadImage).Methods("POST")
	router.HandleFunc("/image/{id}", controllers.GetImageById).Methods("GET")
	router.HandleFunc("/getAllMessages", controllers.GetAllMessages).Methods("GET")
}
