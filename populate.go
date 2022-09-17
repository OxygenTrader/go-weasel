package weasel

func Populate(text string, rate float64, count int) []string {
	var textCopies []string
	for i := 0; i < count; i++ {
		mutatedText := Mutate(text, rate)
		textCopies = append(textCopies, mutatedText)
	}

	return textCopies
}
