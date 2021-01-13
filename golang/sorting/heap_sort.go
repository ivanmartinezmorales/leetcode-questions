package sorting

func HeapSort(array []int) []int {
	return heapSort(array, len(array))
}

func heapSort(array []int, n int) []int {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}

	for i := n - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, i, 0)
	}

	return array
}

func heapify(array []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && array[l] > array[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, n, largest)
	}
}
