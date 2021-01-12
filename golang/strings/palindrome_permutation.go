package main

// CanPermutatePalindrome determines if a permutation of a given string s
// could form a palindrome
func CanPermutatePalindrome(s string) bool {
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, exists := m[s[i]]; !exists {
			m[s[i]] = 1
		} else {
			m[s[i]] += 1
		}
	}

	c := 0
	for k := range m {
		c += m[k] % 2
	}

	return c <= 1
}
