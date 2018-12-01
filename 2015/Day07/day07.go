package Day07

import (
	"strconv"
	"strings"

	"adventofcode/util"
)

type circuit map[string]*node

func (c *circuit) generate() {
	input := util.ReadInputFromFile("input.txt")

	for _, line := range input {
		arr := strings.Split(line, "->")
		arr2 := strings.Split(strings.Trim(arr[0], " "), " ")
		key := strings.Trim(arr[1], " ")
		switch len(arr2) {
		case 1:
			i, err := strconv.Atoi(arr2[0])
			if err != nil {
				(*c)[key] = &node{
					elem1: arr2[0],
					op:    "SET",
				}
			} else {
				(*c)[key] = &node{
					value:  uint32(i),
					cached: true,
				}
			}
		case 2:
			(*c)[key] = &node{
				elem1: arr2[1],
				op:    arr2[0],
			}
		case 3:
			(*c)[key] = &node{
				elem1: arr2[0],
				elem2: arr2[2],
				op:    arr2[1],
			}
		}

		(*c)[key].circ = c
	}
}

type node struct {
	circ   *circuit
	elem1  string
	elem2  string
	op     string
	value  uint32
	cached bool
}

func (n *node) setVal(val uint32) {
	n.value = val
	n.cached = true
}

func (n *node) getVal() uint32 {
	if n.cached {
		return n.value
	}

	var el1, el2 uint32
	if n.elem1 != "" {
		e, err := strconv.Atoi(n.elem1)
		if err != nil {
			el1 = (*n.circ)[n.elem1].getVal()
		} else {
			el1 = uint32(e)
		}
	}

	if n.elem2 != "" {
		e, err := strconv.Atoi(n.elem2)
		if err != nil {
			el2 = (*n.circ)[n.elem2].getVal()
		} else {
			el2 = uint32(e)
		}
	}

	switch n.op {
	case "SET":
		n.setVal(el1)
	case "NOT":
		n.setVal(^el1)
	case "AND":
		n.setVal(el1 & el2)
	case "OR":
		n.setVal(el1 | el2)
	case "LSHIFT":
		n.setVal(el1 << el2)
	case "RSHIFT":
		n.setVal(el1 >> el2)
	default:
		panic("Invalid operand!")
	}

	return n.value
}
