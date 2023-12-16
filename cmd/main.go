package main

import (
	"go-new/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Println("starting server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
