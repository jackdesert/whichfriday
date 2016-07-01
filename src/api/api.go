package main

import (
	"github.com/jackdesert/calendar/src/api/handler"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.Calendar)

	log.Println("Listening1...")
	http.ListenAndServe(":3100", nil)

}
