package codewars

func FindOdd(seq []int) int {
	numberFrequencyMap := map[int]int{}
	for _, num := range seq {
		prev, ok := numberFrequencyMap[num]
		if !ok {
			numberFrequencyMap[num] = 1
		} else {
			numberFrequencyMap[num] = prev + 1
		}
	}
	var res int
	for k, v := range numberFrequencyMap {
		if v%2 == 1 {
			res = k
			break
		}
	}
	return res
}
