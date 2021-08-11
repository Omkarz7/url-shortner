package main

import (
	"fmt"
	"github.com/Omkarz7/url-shortner/api"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cutShort", api.Shorten)
}
