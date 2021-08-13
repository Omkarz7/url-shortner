package api

// cutShortURL srtuct is used to consume request and write reponse
type cutShortURL struct {
	LongURL      string `json:"longURL"`
	ShortenedURL string `json:"shortenedUrl"`
}

//Configs for host and port
var host = "http://localhost"
var Port = ":4700"
