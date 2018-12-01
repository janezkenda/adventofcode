package Day03

type pos struct {
	x int
	y int
}

func getTotalHouses(input string, robo bool) int {
	var p, rp pos

	grid := map[pos]int{
		p: 1,
	}

	for i, c := range input {
		var tempPos *pos
		if i%2 != 0 && robo {
			tempPos = &rp
		} else {
			tempPos = &p
		}

		switch string(c) {
		case "^":
			tempPos.y++
		case "v":
			tempPos.y--
		case "<":
			tempPos.x--
		case ">":
			tempPos.x++
		}

		if _, ok := grid[*tempPos]; !ok {
			grid[*tempPos] = 0
		}

		grid[*tempPos]++
	}

	total := 0
	for _, v := range grid {
		if v > 0 {
			total++
		}
	}

	return total
}
