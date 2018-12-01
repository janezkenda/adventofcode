package Day07

import (
	"fmt"
	"testing"
)

func Test_Day07_Part1(t *testing.T) {
	c := circuit{}
	c.generate()
	fmt.Println(c["a"].getVal())
}

func Test_Day07_Part2(t *testing.T) {
	c := circuit{}
	c.generate()
	c["b"].setVal(956)
	fmt.Println(c["a"].getVal())
}
