package Day10

import (
	"fmt"
	"testing"
)

const input int = 1113222113

/*
1 becomes 11 (1 copy of digit 1).
11 becomes 21 (2 copies of digit 1).
21 becomes 1211 (one 2 followed by one 1).
1211 becomes 111221 (one 1, one 2, and two 1s).
111221 becomes 312211 (three 1s, two 2s, and one 1).
*/

func Test_Day10_Part1(t *testing.T) {
	fmt.Println(lookAndSay(1))
}

func Test_Day10_Part2(t *testing.T) {

}
