package Day09

import (
	"strconv"
	"strings"
	"testing"

	"adventofcode/util"
)

func Test_Day09_Part1(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	dm := distMap{}

	for _, line := range input {
		split := strings.Split(line, " ")
		city1 := split[0]
		city2 := split[2]
		dist, _ := strconv.Atoi(split[4])

		dm.addDistance(city1, city2, dist)
	}

	t.Log(dm)
}

func Test_Day09_Part2(t *testing.T) {

}
