package Day06

import (
	"testing"

	"adventofcode/util"
)

func Test_Day06_Part1(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	g := grid{}

	for _, line := range input {
		g.set(parseLine(line))
	}

	t.Log(g.count())
}

func Test_Day06_Part2(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	ig := intGrid{}

	for _, line := range input {
		ig.set(parseLine(line))
	}

	t.Log(ig.count())
}
