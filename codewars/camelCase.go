package codewars

import "strings"

func CamelCase(s string) string {
	result := ""
	if s == "" {
		return ""
	}
	for _, chunk := range strings.Split(strings.TrimSpace(s), " ") {
		result += strings.ToUpper(string(chunk[0])) + chunk[1:]
	}

	return result
}
