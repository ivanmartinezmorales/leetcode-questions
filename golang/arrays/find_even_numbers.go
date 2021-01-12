package arrays

func findNumbers(nums []int) int {
	v := 0
	for _, k := range nums {
		if isEvenDigits(k) {
			v++
		}
	}
	return v
}

func isEvenDigits(e int) bool {
	c := 0
	for e != 0 {
		e /= 10
		c++
	}
	return c%2 == 0
}
