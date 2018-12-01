package Day02

import (
	"sort"
	"strconv"
	"strings"

	"adventofcode/util"
)

type dim struct {
	length, width, height int
}

func prepareData() []*dim {
	input := util.ReadInputFromFile("input.txt")

	dims := make([]*dim, 0)

	for _, line := range input {
		vars := strings.Split(line, "x")
		numVars := make([]int, 0)
		for _, str := range vars {
			i, _ := strconv.Atoi(str)
			numVars = append(numVars, i)
		}
		dims = append(dims, &dim{
			length: numVars[0],
			width:  numVars[1],
			height: numVars[2],
		})
	}

	return dims
}

func (d *dim) getPaperAmt() int {
	sides := []int{
		d.length * d.width,
		d.width * d.height,
		d.height * d.length,
	}

	m := min(sides)

	for _, i := range sides {
		m += i * 2
	}

	return m
}

func (d *dim) getRibbonAmt() int {
	perimeters := []int{
		d.length + d.width,
		d.width + d.height,
		d.height + d.length,
	}
	bow := d.length * d.width * d.height

	return bow + 2*min(perimeters)
}

func min(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}
