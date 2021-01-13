package sorting

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	sorted := make([]int, len(left)+len(right))
	k, i, j := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sorted[k] = left[i]
			i++
		} else {
			sorted[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		sorted[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		sorted[k] = right[j]
		j++
		k++
	}
	return sorted
}
