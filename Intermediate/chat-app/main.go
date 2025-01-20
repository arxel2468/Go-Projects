package main

import (
	"log"
	"net/http"
	"text/template"

	"chat-app/handlers"

	// "github.com/gorilla/websocket"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/chat.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	})

	http.HandleFunc("/ws", handlers.HandleConnections)

	log.Println("Chat server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
