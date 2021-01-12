package arrays

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	r := []int{}
	for _, k := range nums {
		r = append(r, k*k)
	}
	sort.Ints(r)
	return r
}
