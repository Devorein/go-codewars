package codewars

import (
	"strings"
	"unicode"
)

func ToNato1(words string) string {
	nato := map[string]string{"A": "Alfa", "B": "Bravo", "C": "Charlie", "D": "Delta", "E": "Echo", "F": "Foxtrot", "G": "Golf", "H": "Hotel", "I": "India", "J": "Juliett", "K": "Kilo", "L": "Lima", "M": "Mike", "N": "November", "O": "Oscar", "P": "Papa", "Q": "Quebec", "R": "Romeo", "S": "Sierra", "T": "Tango", "U": "Uniform", "V": "Victor", "W": "Whiskey", "X": "X-ray", "Y": "Yankee", "Z": "Zulu"}
	result := []string{}
	for _, word := range strings.Split(strings.ToUpper(words), " ") {
		for _, letter := range word {
			w, ok := nato[string(letter)]
			if ok {
				result = append(result, w)
			} else {
				result = append(result, string(letter))
			}
		}
	}

	return strings.Join(result, " ")
}

func ToNato2(words string) string {
	nato := []string{
		"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot",
		"Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
		"Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
		"Whiskey", "Xray", "Yankee", "Zulu",
	}

	charToCharlie := map[rune]string{}
	for _, c := range nato {
		charToCharlie[rune(c[0])] = c
	}
	result := ""
	for _, letter := range words {
		if unicode.IsLetter(letter) {
			result += charToCharlie[unicode.ToUpper(letter)] + " "
		} else if unicode.IsPunct(letter) {
			result += string(letter)
		}
	}
	return strings.TrimSpace(result)
}
