package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortURLGeneratorWithCorrectURL(t *testing.T) {
	actual_shortURL := "http://localhost:1337/MTM5MzQ3"
	actual_longURL := "https://www.yahoo.com"

	StoreURL(actual_shortURL, actual_longURL)

	expected_shortURL := FindShortURL(actual_longURL)
	expected_longURL := FindLongURL(actual_shortURL)

	assert.Equal(t, actual_shortURL, expected_shortURL)
	assert.Equal(t, actual_longURL, expected_longURL)
}

func TestShortURLGeneratorWithWrongURL(t *testing.T) {
	actual_shortURL := "http://localhost:1337/MTM5MzQ3"
	actual_longURL := "https://www.yahoo.com"

	StoreURL(actual_shortURL, actual_longURL)

	expected_shortURL := FindShortURL("absent")
	expected_longURL := FindLongURL("absent")

	assert.Equal(t, "", expected_shortURL)
	assert.Equal(t, "", expected_longURL)
}
