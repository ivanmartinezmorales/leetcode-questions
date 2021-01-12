package arrays

func HasSingleCycle(array []int) bool {
	return hasSingleCycle(array, 0)
}

func hasSingleCycle(array []int, startingIndex int) bool {
	// O(n) time | O(1) space
	elementsVisited := 0
	current := startingIndex

	for elementsVisited < len(array) {
		if elementsVisited > 0 && current == startingIndex {
			return false
		}
		elementsVisited++
		current = getNextIdx(current, array)
	}

	return current == startingIndex
}

func getNextIdx(current int, array []int) int {
	jump := array[current]

	if nextIdx := (current + jump) % len(array); nextIdx >= 0 {
		return nextIdx
	} else {
		return nextIdx + len(array)
	}
}
