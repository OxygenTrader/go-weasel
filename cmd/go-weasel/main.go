package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	weasel "github.com/OxygenTrader/go-weasel"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sample := flag.String("sample", "METHINKS IT IS LIKE A WEASEL", "target string")
	rate := flag.Float64("rate", 0.05, "character mutate rate")
	populationCount := flag.Int("count", 100, "population size")
	flag.Parse()

	var generationCounter int
	currentVariant := weasel.Initialize(len(*sample))
	for currentVariant != *sample {
		const outputRate = 10
		if generationCounter%outputRate == 0 {
			fmt.Println(generationCounter, currentVariant)
		}

		variants := weasel.Populate(currentVariant, *rate, *populationCount)
		currentVariant = weasel.Search(variants, *sample)

		generationCounter++
	}

	fmt.Println(generationCounter, currentVariant)
}
