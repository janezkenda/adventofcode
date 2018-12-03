package Day02

import (
	"testing"

	"adventofcode/util"
)

func Test_Day02_Part1(t *testing.T) {
	double, tripple := 0, 0
	input := util.ReadInputFromFile("input.txt")

	for _, line := range input {
		lc := NewLetterCount(line)
		if lc.HasDouble() {
			double++
		}
		if lc.HasTripple() {
			tripple++
		}
	}

	t.Log(double * tripple)
}

func Test_Day02_Part2(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")
	t.Log(FindSimilarIDs(input))
}
