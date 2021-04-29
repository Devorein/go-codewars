package codewars

func DeadFishSwim(data string) []int {
	val := 0
	values := []int{}

	for _, char := range data {
		switch string(char) {
		case "i":
			val += 1
		case "d":
			val -= 1
		case "s":
			val *= val
		case "o":
			values = append(values, val)
		}
	}
	return values
}
