package arrays

func findKthPositive(arr []int, k int) int {
	if k <= arr[0]-1 {
		return k
	}

	k -= arr[0] - 1

	n := len(arr)

	for i := 0; i < n-1; i++ {
		c := arr[i+1] - arr[i] - 1
		if k <= c {
			return arr[i] + k
		}

		k -= c
	}

	return arr[n-1] + k
}
