package codewars

import "fmt"

func PrinterError(s string) string {
	counter := 0

	for _, v := range s {
		if v > 109 {
			counter++
		}
	}
	return fmt.Sprintf("%v/%v", counter, len(s))
}
