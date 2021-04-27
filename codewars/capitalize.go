package codewars

import "strings"

func CapitalizeEvenOdd(st string) []string {
	evenCapitalize, oddCapitalize := "", ""
	for i, v := range st {
		if i%2 == 1 {
			oddCapitalize += strings.ToUpper(string(v))
			evenCapitalize += string(v)
		} else {
			evenCapitalize += strings.ToUpper(string(v))
			oddCapitalize += string(v)
		}
	}

	return []string{evenCapitalize, oddCapitalize}
}
