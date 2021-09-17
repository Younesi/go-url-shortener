package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://virgool.io/@younesi"
	shortLink_1 := GenerateShortLink(initialLink_1)

	initialLink_2 := "https://virgool.io/@younesi/array-and-slices-in-golang-l6ye1tbzuckj"
	shortLink_2 := GenerateShortLink(initialLink_2)

	initialLink_3 := "https://virgool.io/golangpub/structs-in-golang-lmedjf9xijyz"
	shortLink_3 := GenerateShortLink(initialLink_3)

	assert.Equal(t, shortLink_1, "WvCw4bw9")
	assert.Equal(t, shortLink_2, "6N8gWFjQ")
	assert.Equal(t, shortLink_3, "LMEZJZWS")
}
