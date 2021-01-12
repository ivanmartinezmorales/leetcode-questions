package string

import (
	"strings"
)

func ArrayStringsAreEqual(w1, w2 []string) bool {
	return strings.Join(w1, "") == strings.Join(w2, "")
}
