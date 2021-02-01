package arrays

func MissingNumber(nums []int) int {
	var expected int = len(nums) * (len(nums) + 1) / 2
	actual := sum(nums)
	return expected - actual
}

func sum(a []int) int {
	r := 0
	for _, n := range a {
		r += n
	}
	return r
}
