package Day05

import (
	"fmt"
	"testing"

	"adventofcode/util"
)

func Test_Day05_Part1(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	count := 0

	for _, line := range input {
		if isNice1(line) {
			count++
		}
	}

	fmt.Println(count)
}

func Test_Day05_Part2(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	count := 0

	for _, line := range input {
		if isNice2(line) {
			count++
		}
	}

	fmt.Println(count)
}
