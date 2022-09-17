package weasel

func Fitness(text string, sample string) int {
	var count int
	for index, character := range text {
		if character != rune(sample[index]) {
			count++
		}
	}

	return count
}
