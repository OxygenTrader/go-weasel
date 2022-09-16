package weasel

import (
	"math/rand"
	"testing"
)

func TestMutate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	actualMutatedText := Mutate(text, 0.2)

	expectedMutatedText := "the qu Pk brown fox jumps oveF tGD Nazy dog"
	if actualMutatedText != expectedMutatedText {
		test.Fail()
	}
}
