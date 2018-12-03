package Day02

import "bytes"

type lettercount map[int32]int

func NewLetterCount(word string) *lettercount {
	lc := lettercount{}
	for _, b := range word {
		lc[b]++
	}
	return &lc
}

func (lc *lettercount) HasDouble() bool {
	for _, c := range *lc {
		if c == 2 {
			return true
		}
	}
	return false
}

func (lc *lettercount) HasTripple() bool {
	for _, c := range *lc {
		if c == 3 {
			return true
		}
	}
	return false
}

func FindChecksum(input []string) int {
	double, tripple := 0, 0

	for _, line := range input {
		lc := NewLetterCount(line)
		if lc.HasDouble() {
			double++
		}
		if lc.HasTripple() {
			tripple++
		}
	}

	return double * tripple
}

func FindSimilarIDs(input []string) string {
	for _, line := range input {
		for _, line2 := range input {
			if line == line2 {
				continue
			}
			match := CheckIfMatching(line, line2)

			if match != "" {
				return match
			}
		}
	}

	return ""
}

func CheckIfMatching(word1, word2 string) string {
	var buf bytes.Buffer
	differences := 0

	for i, c := range word1 {
		if string(word2[i]) == string(c) {
			buf.WriteString(string(c))
		} else {
			differences++
			if differences > 1 {
				return ""
			}
		}
	}

	return buf.String()
}
