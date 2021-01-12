package main

func SunsetViews(b []int, d string) []int {
	// Given an array of buildings and a direction,
	// Return an array of buildings that can see the sunset

	// A building can see the sunset if its strictly taller than
	// All of the buildings that come after it in the direction
	// that it faces.

	// Input array named buildings contains positive, non-zero integers
	// the building at index i thus has a height denoted by buildings[i]

	// All of the buildings face in the same direction, and this
	// direction can either be east or west denoted by the input string
	// named direction, which is always going to be equal to "EAST" or "WEST"

	// Indices should be sorted in ascending order

	return []int{}
}

func reverse(a []int) {
	l := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		a[i], a[l-i] = a[l-i], a[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
