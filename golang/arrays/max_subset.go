package arrays

func MaxSubsetSumNoAdjacent(a []int) int {
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 {
		return a[0]
	}

	f := max(a[0], a[1])
	s := a[0]

	for i := 2; i < len(a); i++ {
		cur := max(f, s+a[i])
		s = f
		f = cur
	}
	return f
}

func MaxSubsetSumNoAdjacentNaive(a []int) int {
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 {
		return a[0]
	}

	m := make([]int, len(a))
	m[0], m[1] = a[0], max(a[0], a[1])

	for i := 2; i < len(a); i++ {
		m[i] = max(m[i-1], m[i-2]+a[i])
	}

	return m[len(a)-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
