package string

type substring struct {
	left  int
	right int
}

func (ss substring) length() int {
	return ss.right - ss.left
}

func LongestPalindromicSubstring(str string) string {
	result := substring{0, 1}
	for i := 1; i < len(str); i++ {
		odd := getLongestPalindrome(str, i-1, i+1)
		even := getLongestPalindrome(str, i-1, i)

		longest := even
		if odd.length() > result.length() {
			result = odd
		}

		if longest.length() > result.length() {
			result = longest
		}
	}
	return str[result.left:result.right]
}

func getLongestPalindrome(str string, left, right int) substring {
	for left >= 0 && right < len(str) {
		if str[left] != str[right] {
			break
		}
		left -= 1
		right += 1
	}
	return substring{left + 1, right}
}
