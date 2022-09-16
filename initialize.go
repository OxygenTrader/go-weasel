package weasel

import "math/rand"

func makeRandomCharacter() rune {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	randomIndex := rand.Intn(len(alphabet))
	return rune(alphabet[randomIndex])
}

func Initialize(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result = result + string(makeRandomCharacter())
	}

	return result
}
