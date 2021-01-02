package main

func FirstDuplicateValue(array []int) int {

	// O(n) time | O(n) space
	seen := map[int]bool{}

	for _, value := range array {
		if seen[value] {
			return value
		} else {
			seen[value] = true
		}
	}
	return -1
}
