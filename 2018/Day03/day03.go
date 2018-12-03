package Day03

import (
	"strconv"
	"strings"
)

type (
	pos struct {
		x int
		y int
	}

	Area struct {
		area    map[pos]int
		overlap map[int]struct{}
	}
)

func NewArea() *Area {
	return &Area{
		area:    map[pos]int{},
		overlap: map[int]struct{}{},
	}
}

func (a *Area) Insert(start pos, width, height, id int) {
	for i := start.x; i < start.x+width; i++ {
		for j := start.y; j < start.y+height; j++ {
			currentPos := pos{
				i,
				j,
			}

			if v, ok := (*a).area[currentPos]; ok {
				(*a).area[currentPos] = -1
				(*a).overlap[v] = struct{}{}
				(*a).overlap[id] = struct{}{}
			} else {
				(*a).area[currentPos] = id
			}
		}
	}
}

func (a *Area) CountOverlap() int {
	count := 0
	for _, v := range (*a).area {
		if v < 0 {
			count++
		}
	}

	return count
}

func (a *Area) FindNonOverlap(input []string) int {
	for _, l := range input {
		id, _ := strconv.Atoi(strings.Split(strings.Trim(l, "#"), " ")[0])
		if _, ok := (*a).overlap[id]; !ok {
			return id
		}
	}

	return -1
}
