package weasel

import "testing"

func TestSearch_withMatch(test *testing.T) {
	variants := []string{"cat", "dog", "pig"}
	sample := "dig"
	actualResult := Search(variants, sample)

	expectedResult := "dog"
	if actualResult != expectedResult {
		test.Fail()
	}
}

func TestSearch_withNoMatch(test *testing.T) {
	variants := []string{"cat", "dog", "pig"}
	sample := "one"
	actualResult := Search(variants, sample)

	expectedResult := "cat"
	if actualResult != expectedResult {
		test.Fail()
	}
}
