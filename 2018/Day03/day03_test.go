package Day03

import (
	"strconv"
	"strings"
	"testing"

	"adventofcode/util"
)

func Test_Day03_Part1(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	a := NewArea()

	for _, line := range input {
		arr := strings.Split(line, " ")
		id, _ := strconv.Atoi(arr[0][1:])
		sc := strings.Split(strings.Trim(arr[2], ":"), ",")
		xpos, _ := strconv.Atoi(sc[0])
		ypos, _ := strconv.Atoi(sc[1])
		dim := strings.Split(arr[3], "x")
		w, _ := strconv.Atoi(dim[0])
		h, _ := strconv.Atoi(dim[1])
		p := pos{
			x: xpos,
			y: ypos,
		}
		a.Insert(p, w, h, id)
	}

	t.Log(a.CountOverlap())
	t.Log(a.FindNonOverlap(input))
}

func Test_Day03_Part2(t *testing.T) {
	input := util.ReadInputFromFile("input.txt")

	_ = input
}
