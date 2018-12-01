package Day02

import "testing"

func Test_Day02_Part1(t *testing.T) {
	data := prepareData()

	totalPaper := 0
	for _, d := range data {
		totalPaper += d.getPaperAmt()
	}
	t.Log(totalPaper)
}

func Test_Day02_Part2(t *testing.T) {
	data := prepareData()

	totalRibbon := 0
	for _, d := range data {
		totalRibbon += d.getRibbonAmt()
	}
	t.Log(totalRibbon)
}
