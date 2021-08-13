package shorty

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

// CreateShortLink hashes a given URL into a shorter one
func CreateShortLink(initialLink string) string {
	urlHashBytes := sha256.Sum256([]byte(initialLink + uuid.New().String()))
	generatedNumber := new(big.Int).SetBytes(urlHashBytes[:]).Uint64()
	finalString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(generatedNumber)))
	return `/` + finalString[:8]
}

// CreateShortLinkAlternate is and alternate way to shorten URL
func CreateShortLinkAlternate(initialLink string) string {
	return uuid.New().String()[:8]
}
