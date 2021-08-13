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

// HomePagePing is a dummy URL test if server is up
func HomePagePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

// PageNotFound is used to handle incorrect URL
func PageNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

// RedirectShortURL redirects shortURL to longURL, handles error  otherwise
func RedirectShortURL(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/" {
		HomePagePing(w, r)
		return
	}
	longURL := store.FindLongURL(r.RequestURI)
	if longURL == "" {
		PageNotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusSeeOther)
}
