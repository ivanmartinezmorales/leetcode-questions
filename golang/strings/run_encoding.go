package string

import (
	"strconv"
)

func RunLengthEncoding(s string) string {
	c := []byte{}
	r := 1

	for i := 1; i < len(s); i++ {
		cur := s[i]
		prev := s[i-1]
		if cur != prev || r == 9 {
			c = append(c, strconv.Itoa(c)[0])
			c = append(c, prev)
			r = 0
		}
		r++
	}

	c = append(c, strconv.Itoa(c)[0])
	c = append(c, s[len(s)-1])

	return string(c)
}
