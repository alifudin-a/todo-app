package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const aiueo = "aiueo"

// CreateRandomTitle will generate a random string by n, max length by n
func createRandomTitle(n int) string {
	var sb strings.Builder
	k := len(aiueo)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func createRandomDescription(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomTitle generate random title
func RandomTitle() string {
	return createRandomTitle(5)
}

// RandomDescription generate random description
func RandomDescription() string {
	return createRandomDescription(30)
}
