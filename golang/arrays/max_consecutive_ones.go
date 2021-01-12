package arrays

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0

	for _, item := range nums {
		if item == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if max > count {
		return max
	}
	return count
}
