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
	reached := map[int]struct{}{
		0: struct{}{},
	}

	return getFirstFrequencyReq(input, 0, reached)
}

func getFirstFrequencyReq(input []string, sum int, reached map[int]struct{}) int {
	for _, i := range input {
		num, _ := strconv.Atoi(i)
		sum += num
		if _, ok := reached[sum]; ok {
			return sum
		}
		reached[sum] = struct{}{}
	}

	return getFirstFrequencyReq(input, sum, reached)
}
