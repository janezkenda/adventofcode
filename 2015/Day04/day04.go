package Day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func findFirst(key, subs string) int {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", key, i)))
		if strings.Index(hex.EncodeToString(hash[:]), subs) == 0 {
			return i
		}
	}
}
