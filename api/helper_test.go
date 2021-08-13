package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortURLGenerator(t *testing.T) {
	initialLink_1 := "https://www.google.com"
	shortLink_1 := GenerateShortURL(initialLink_1)

	shortLink_2 := GenerateShortURL(initialLink_1)

	initialLink_2 := "https://www.yahoo.com"
	shortLink_3 := GenerateShortURL(initialLink_2)

	assert.Equal(t, shortLink_1, shortLink_2)
	assert.NotEqual(t, shortLink_3, shortLink_1)
	assert.NotEqual(t, shortLink_3, shortLink_2)
}
