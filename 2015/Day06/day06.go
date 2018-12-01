package Day06

import (
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

type grid map[pos]bool

func (g *grid) set(min, max pos, action string) {
	for i := min.x; i <= max.x; i++ {
		for j := min.y; j <= max.y; j++ {
			key := pos{
				x: i,
				y: j,
			}

			switch action {
			case "on":
				(*g)[key] = true
			case "off":
				(*g)[key] = false
			case "toggle":
				if val, ok := (*g)[key]; ok {
					(*g)[key] = !val
				} else {
					(*g)[key] = true
				}
			}
		}
	}
}

func (g *grid) count() int {
	count := 0
	for _, g := range *g {
		if g {
			count++
		}
	}

	return count
}

type intGrid map[pos]int

func (ig *intGrid) set(min, max pos, action string) {
	for i := min.x; i <= max.x; i++ {
		for j := min.y; j <= max.y; j++ {
			key := pos{
				x: i,
				y: j,
			}
			if _, ok := (*ig)[key]; !ok {
				(*ig)[key] = 0
			}

			switch action {
			case "on":
				(*ig)[key]++
			case "off":
				(*ig)[key]--
			case "toggle":
				(*ig)[key] += 2
			}

			if (*ig)[key] < 0 {
				(*ig)[key] = 0
			}
		}
	}
}

func (ig *intGrid) count() int {
	count := 0
	for _, g := range *ig {
		count += g
	}

	return count
}

func parseLine(line string) (pos, pos, string) {
	arr := strings.Split(line, " ")
	return parsePos(arr[len(arr)-3]), parsePos(arr[len(arr)-1]), arr[len(arr)-4]
}

func parsePos(str string) pos {
	arr := strings.Split(str, ",")
	return pos{
		x: getInt(arr[0]),
		y: getInt(arr[1]),
	}
}

func getInt(str string) int {
	s, _ := strconv.Atoi(str)
	return s
}
