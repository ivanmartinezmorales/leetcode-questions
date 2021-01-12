package arrays

func LongestPeak(array []int) int {
	if len(array) < 3 {
		return 0
	}

	longest := 0
	i := 1

	for i < len(array)-1 {
		if isPeak := array[i-1] < array[i] && array[i] > array[i+1]; !isPeak {
			i++
			continue
		}

		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
			leftIdx--
		}

		rightIdx := i + 2
		for rightIdx < len(array) && array[rightIdx] < array[rightIdx-1] {
			rightIdx++
		}

		if current := rightIdx - leftIdx - 1; current > longest {
			longest = current
		}

		i = rightIdx
	}

	return longest
}
