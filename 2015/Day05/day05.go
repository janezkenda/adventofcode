package Day05

import (
	"fmt"
	"strings"
)

func isNice1(word string) bool {
	for _, comb := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Index(word, comb) > -1 {
			return false
		}
	}

	vowels := 0
	for _, vowel := range "aeiou" {
		vowels += strings.Count(word, string(vowel))
	}

	if vowels < 3 {
		return false
	}

	prevChar := string(word[0])

	for _, c := range word[1:] {
		if string(c) == prevChar {
			return true
		}
		prevChar = string(c)
	}

	return false
}

func isNice2(word string) bool {
	var got bool
	for c := 'a'; c <= 'z' && got == false; c++ {
		for c2 := 'a'; c2 <= 'z' && got == false; c2++ {
			if strings.Count(word, fmt.Sprintf("%c%c", c, c2)) >= 2 {
				got = true
			}
		}
	}

	if !got {
		return false
	}

	for i, c := range word[2:] {
		if string(word[i]) == string(c) {
			return true
		}
	}

	return false
}
