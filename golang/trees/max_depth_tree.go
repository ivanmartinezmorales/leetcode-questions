package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(r *TreeNode) int {
	if r == nil {
		return 0
	}
	q := []*TreeNode{r}
	d := 1
	s, e := 0, 1

	for {
		nC := 0
		for i := s; i < e; i++ {
			n := q[i]
			if n.Left != nil {
				q = append(q, n.Left)
				nC++
			}

			if n.Right != nil {
				q = append(q, n.Right)
				nC++
			}
		}

		if nC == 0 {
			break
		}
		d++
		s = e
		e = e + nC
	}

	return d
}
