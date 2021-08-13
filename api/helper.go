package api

import (
	"github.com/Omkarz7/url-shortner/shorty"
	"github.com/Omkarz7/url-shortner/store"
)

//GenerateShortURL returns a shortened URL for a given long URL
func GenerateShortURL(longURL string) string {
	shortURL := store.FindShortURL(longURL)
	if shortURL == "" {
		shortURL = shorty.CreateShortLink(longURL)
		store.StoreURL(shortURL, longURL)
	}
	shortURL = host + Port + shortURL
	return shortURL
}
