package weasel

import "testing"

func TestFitness_withDifferences(test *testing.T) {
	text := "the quick brown fox jumps over the lazy dog"
	sample := "the quick brown cat jumps over the lazy pig"

	actualCount := Fitness(text, sample)

	expectedCount := 5
	if actualCount != expectedCount {
		test.Fail()
	}
}

func TestFitness_withNoDifferences(test *testing.T) {
	text := "the quick brown fox jumps over the lazy dog"

	actualCount := Fitness(text, text)

	expectedCount := 0
	if actualCount != expectedCount {
		test.Fail()
	}
}
