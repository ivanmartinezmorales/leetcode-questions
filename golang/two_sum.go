package main

import "sort"

// TwoNumberSum finds two numbers in an array of integers that equal a target sum
func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum > target {
			right = right - 1
		} else {
			left = left + 1
		}
	}
	return nil
}
