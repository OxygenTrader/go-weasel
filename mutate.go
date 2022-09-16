package weasel

import "math/rand"

func Mutate(text string, rate float64) string {
	var mutatedText string
	for _, character := range text {
		if rand.Float64() < rate {
			mutatedText = mutatedText + string(makeRandomCharacter())
		} else {
			mutatedText = mutatedText + string(character)
		}
	}

	return mutatedText
}
