package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Omkarz7/url-shortner/store"
)

// GenerateShortURL handle cutShort route to shorten URLs
func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	var convertion cutShortURL
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &convertion)

	convertion.ShortenedURL = GenerateShortURL(convertion.LongURL)
	json.NewEncoder(w).Encode(convertion)
}
