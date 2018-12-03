package Day02

import (
	"testing"

	"adventofcode/util"
)

func Test_Day02_Part1(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")
	t.Log(FindChecksum(input))
}

func Test_Day02_Part2(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")
	t.Log(FindSimilarIDs(input))
}
