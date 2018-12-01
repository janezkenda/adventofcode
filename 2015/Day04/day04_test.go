package Day04

import "testing"

const input string = "bgvyzdsv"

func Test_Day03_Part1(t *testing.T) {
	t.Log(findFirst(input, "00000"))
}

func Test_Day03_Part2(t *testing.T) {
	t.Log(findFirst(input, "000000"))
}
