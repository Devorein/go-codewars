package codewars

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	tuples := []Tuple{}
	charFrequencyMap := map[rune]int{}

	for _, v := range text {
		charFrequencyMap[rune(v)]++
	}

	for k, v := range charFrequencyMap {
		tuples = append(tuples, Tuple{Char: rune(k), Count: v})
	}

	return tuples
}
