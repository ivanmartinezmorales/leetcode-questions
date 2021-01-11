package utils

import (
	"math"
)

// Contains returns true if the slice contains the value
func Contains(a []int, b int) bool {
	for _, k := range a {
		if k == b {
			return true
		}
	}
	return false
}

// MaxInt gets the margest int in the bunch
func MaxInt(a ...int) int {
	switch len(a) {
	case 0:
		return 0
	case 1:
		return a[0]
	case 2:
		if a[0] > a[1] {
			return a[0]
		} else {
			return a[1]
		}
	default:
		m := math.MinInt32
		for _, k := range a {
			if k > m {
				m = k
			}
		}
		return m
	}
}

// MinInt returns the smallest int in the bunch.
func MinInt(a ...int) int {
	switch len(a) {
	case 0:
		return 0
	case 1:
		return a[0]
	case 2:
		if a[0] > a[1] {
			return a[1]
		} else {
			return a[0]
		}
	default:
		m := math.MaxInt32
		for _, k := range a {
			if k < m {
				m = k
			}
		}
		return m
	}
}
