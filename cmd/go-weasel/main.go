package main

import (
	"fmt"
	"math/rand"
	"time"

	weasel "github.com/OxygenTrader/go-weasel"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sample := "METHINKS IT IS LIKE A WEASEL"
	var generationCounter int
	currentVariant := weasel.Initialize(len(sample))
	for currentVariant != sample {
		const outputRate = 10
		if generationCounter%outputRate == 0 {
			fmt.Println(generationCounter, currentVariant)
		}

		variants := weasel.Populate(currentVariant, 0.05, 100)
		currentVariant = weasel.Search(variants, sample)

		generationCounter++
	}

	fmt.Println(generationCounter, currentVariant)
}
