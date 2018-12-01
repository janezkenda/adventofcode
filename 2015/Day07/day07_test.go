package Day07

import "testing"

func Test_Day07_Part1(t *testing.T) {
	c := circuit{}
	c.generate()
	t.Log(c["a"].getVal())
}

func Test_Day07_Part2(t *testing.T) {
	c := circuit{}
	c.generate()
	c["b"].setVal(956)
	t.Log(c["a"].getVal())
}
