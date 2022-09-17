package weasel

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPopulate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	actualTextCopies := Populate(text, 0.2, 3)

	expectedTextCopies := []string{
		"the qu Pk brown fox jumps oveF tGD Nazy dog",
		"the N ick broRn foS jumps oveB theTlazyEdQg",
		"the quWck b Ewn foxMjumNs Qver She lRzy dog",
	}
	if !reflect.DeepEqual(actualTextCopies, expectedTextCopies) {
		test.Fail()
	}
}
