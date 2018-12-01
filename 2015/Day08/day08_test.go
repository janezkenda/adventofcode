package Day08

import (
	"testing"

	"adventofcode/util"
)

func Test_Day08_Part1(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	count := 0
	for _, line := range input {
		count += len(line) - unquotedLen(line)
	}

	t.Log(count)
}

func Test_Day08_Part2(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	count := 0
	for _, line := range input {
		count += doubleQuotedLen(line) - len(line)
	}

	t.Log(count)
}
