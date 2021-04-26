package codewars

import (
	"strings"
)

func AbbrevName(name string) string {
	substr := strings.Split(strings.ToUpper(name), " ")
	return string(substr[0][0]) + "." + string(substr[1][0])
}
