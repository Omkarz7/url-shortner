package store

// shortenedURLStore keeps track to URLs already shortened
var shortenedURLStore = make(map[string]string)

// FindShortURL fetches URL from memory if exists
func FindShortURL(url string) string {
	for key, value := range shortenedURLStore {
		if value == url {
			return key
		}
	}
	return ""
}

// StoreURL stores URL in memory
func StoreURL(shortURL, longURL string) {
	shortenedURLStore[shortURL] = longURL
}
