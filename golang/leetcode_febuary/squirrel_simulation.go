package leetcode_febuary

import "math"

func MinDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	t := 0
	d := math.MinInt64
	for _, nut := range nuts {
		t += distance(nut, tree) * 2
		d = max(d, distance(nut, tree)-distance(nut, squirrel))
	}
	return t - d
}

func distance(a, b []int) int {
	x := float64(a[0] - b[0])
	y := float64(a[1] - b[1])
	return int(math.Abs(x) + math.Abs(y))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
