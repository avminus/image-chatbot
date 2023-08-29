package main

import (
	"fmt"
	"image-chatbot/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterChatBotRoutes(r)
	fmt.Printf("Starting the server at port: 8000")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
