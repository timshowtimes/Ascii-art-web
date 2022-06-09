package main

import (
	handlers "ascii-art-web-stylize/pkg"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

func main() {
	home := handlers.Home
	formHandler := handlers.FormHandler
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home) // Execute formhandler when submit button pressed
	mux.HandleFunc("/ascii-art", formHandler)
	fmt.Printf("Starting server - http://localhost:%v\n", PORT)
	if err := http.ListenAndServe(":"+PORT, mux); err != nil { // start the server
		log.Fatal(err)
	}
}
