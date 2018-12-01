package Day08

import "strconv"

func unquotedLen(word string) int {
	u, _ := strconv.Unquote(word)
	return len(u)
}

func doubleQuotedLen(word string) int {
	return len(strconv.Quote(word))
}
