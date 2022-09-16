package weasel

import (
	"math/rand"
	"testing"
)

func TestInitialize(test *testing.T) {
	rand.Seed(1)

	length := 5
	actualResult := Initialize(length)

	expectedResult := "OPCLE"
	if actualResult != expectedResult {
		test.Fail()
	}
}
