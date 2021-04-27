package codewars

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	totalSeconds := float32(g*3600) / float32(v2-v1)
	hours := int(totalSeconds) / 3600
	totalSeconds = float32(int(totalSeconds) % 3600)
	minutes := int(totalSeconds) / 60
	seconds := int(totalSeconds) % 60
	return [3]int{hours, minutes, seconds}
}
