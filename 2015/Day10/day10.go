package Day10

import (
	"fmt"
	"strconv"
)

func lookAndSay(num int) int {
	str := strconv.Itoa(num)

	res := ""
	temp := str[0]
	counter := 0
	for i := 0; i < len(str); i++ {
		counter++
		if str[i] != temp || i == len(str) {
			res = fmt.Sprintf("%s%d%s", res, counter, string(temp))
			temp = str[i]
			counter = 1
		}
	}

	ret, _ := strconv.Atoi(res)

	return ret
}
