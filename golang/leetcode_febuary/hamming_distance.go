package leetcode_febuary

// Write a function that takes an unsigned integer and returns the 1 bits it has.
//
// For my approach, I used 2s complement. This would be a great way to approach
// the problem because I could use the number, divide it by 2, such that if the
// answer was odd, then I would increment the count
func HammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		if num%2 == 1 {
			ans += 1
		}
		num = num / 2
	}
	return ans
}
