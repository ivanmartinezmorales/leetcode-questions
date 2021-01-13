package sorting

// QuickSort sorts an array using the quick sort algorithm. O(n^2) time | O(log(n)) space
func QuickSort(a []int) []int {
	return quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, low, high int) []int {
	if low < high {
		p = partition(a, low, high)

		quickSort(a, low, p-1)
		quickSort(a, p+1, high)
	}
	return a
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i] // swap
		}
	}

	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}
