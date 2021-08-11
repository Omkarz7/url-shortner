package api

// cutShortURL srtuct is used to consume request and write reponse
type cutShortURL struct {
	OriginalURL  string `json:"originalUrl"`
	ShortenedURL string `json:"shortenedUrl"`
}

// shortenedURLMap keeps track to URLs already shortened
var shortenedURLMap = make(map[string]string)
