package Day01

import "strings"

func getFloor(input string) int {
	return strings.Count(input, "(") - strings.Count(input, ")")
}

func getBasement(input string) int {
	pos := 0
	for i, char := range input {
		switch string(char) {
		case "(":
			pos++
		case ")":
			pos--
		}
		if pos == -1 {
			return i + 1
		}
	}

	return pos
}
