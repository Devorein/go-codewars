package codewars

func RowSumOddNumbers(n int) int {
	start := (n * n) - (n - 1)
	res := start
	for i := 1; i < n; i++ {
		res += start + (2 * i)
	}
	return res
}
