package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Omkarz7/url-shortner/api"
)

func main() {
	fmt.Println("Starting server at port ", api.Port)

	http.HandleFunc("/cutShort", api.CreateShortUrl)
	http.HandleFunc("/", api.RedirectShortURL)

	if err := http.ListenAndServe(api.Port, nil); err != nil {
		log.Fatal(err)
	}
}
