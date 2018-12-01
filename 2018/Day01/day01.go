package Day01

import "strconv"

func sumItUp(input []string) int {
	sum := 0

	for _, i := range input {
		num, _ := strconv.Atoi(i)
		sum += num
	}

	return sum
}

func getFirstFrequency(input []string) int {
	sum := 0

	reached := map[int]struct{}{
		0: struct{}{},
	}

	for {
		for _, i := range input {
			num, _ := strconv.Atoi(i)
			sum += num
			if _, ok := reached[sum]; ok {
				return sum
			}
			reached[sum] = struct{}{}
		}
	}
}
