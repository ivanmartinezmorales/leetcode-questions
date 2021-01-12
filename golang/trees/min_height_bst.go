package trees

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func MinHeightBST(array []int) *BST {
	return minHeightBST(array, 0, len(array)-1)
}

func minHeightBST(array []int, start, end int) *BST {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	bst := &BST{Value: array[mid]}
	bst.Left = minHeightBST(array, start, mid-1)
	bst.Right = minHeightBST(array, mid+1, end)
	return bst
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
