package Day01

import (
	"adventofcode/util"
	"testing"
)

func Test_Day01_Part1(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")
	t.Log(sumItUp(input))
}

func Test_Day01_Part2(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")
	t.Log(getFirstFrequency(input))
}
