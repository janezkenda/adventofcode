package Day10

import "testing"

const input = "1113222113"

func Test_Day10_Part1(t *testing.T) {
	temp := input
	for i := 0; i < 40; i++ {
		temp = lookAndSay(temp)
	}

	t.Log(len(temp))
}

func Test_Day10_Part2(t *testing.T) {
	temp := input
	for i := 0; i < 50; i++ {
		temp = lookAndSay(temp)
	}

	t.Log(len(temp))
}
