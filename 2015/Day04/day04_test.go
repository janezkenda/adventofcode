package Day04

import (
	"fmt"
	"testing"
)

const input string = "bgvyzdsv"

func Test_Day03_Part1(t *testing.T) {
	fmt.Println(findFirst(input, "00000"))
}

func Test_Day03_Part2(t *testing.T) {
	fmt.Println(findFirst(input, "000000"))
}
