package weasel

func Search(variants []string, sample string) string {
	var result string
	minCount := len(sample) + 1
	for _, variant := range variants {
		count := Fitness(variant, sample)
		if count < minCount {
			minCount = count
			result = variant
		}
	}

	return result
}
