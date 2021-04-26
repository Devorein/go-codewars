package main

import (
	"fmt"
	"sort"
)

func IsTriangle1(ints ...int) bool {
	sort.Ints(ints)
	if ints[0] <= 0 || ints[1] <= 0 || ints[2] <= 0 {
		return false
	} else if ints[0]+ints[1] > ints[2] {
		return true
	} else {
		return false
	}
}

func findMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func IsTriangle2(a, b, c int) bool {
	longSide := findMax(a, findMax(b, c))
	if a+b+c-longSide > longSide {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsTriangle1(5, 1, 2))
	fmt.Println(IsTriangle2(5, 1, 2))
}
