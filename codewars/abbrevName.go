package codewars

import (
	"fmt"
	"strings"
)

func AbbrevName1(name string) string {
	substr := strings.Split(strings.ToUpper(name), " ")
	return string(substr[0][0]) + "." + string(substr[1][0])
}

func AbbrevName2(name string) string {
	var parts []string
	for _, part := range strings.Split(strings.ToUpper(name), " ") {
		parts = append(parts, part[:1])
	}
	return strings.Join(parts, ".")
}

func AbbrevName3(name string) string {
	return fmt.Sprintf("%c.%c", strings.ToUpper(name)[0], strings.ToUpper(name)[strings.Index(name, " ")+1])
}
