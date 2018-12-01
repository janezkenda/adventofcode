package Day10

import (
	"bytes"
	"strconv"
)

func lookAndSay(input string) string {
	var buf bytes.Buffer
	temp := input[0]
	counter := 0

	for _, char := range []byte(input) {
		if char != temp {
			buf.WriteString(strconv.Itoa(counter))
			buf.WriteByte(temp)
			counter = 0
			temp = char
		}
		counter++
	}

	buf.WriteString(strconv.Itoa(counter))
	buf.WriteByte(temp)

	return buf.String()
}
