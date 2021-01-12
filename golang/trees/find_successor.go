package trees

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func findSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	tOrder := []*BinaryTree{}
	getInOrderTraversal(tree, &traversalOrder)

	for i, current := range tOrder {
		if current != node {
			continue
		}

		if i == len(tOrder)-1 {
			return nil
		}
		return tOrder[i+1]
	}
	return nil
}

func getInOrderTraversal(node *BinaryTree, order *[]*BinaryTree) {
	if node == nil {
		return
	}

	getInOrderTraversal(node.Left, order)
	*order = append(*order, node)
	getInOrderTraversal(node.Right, order)
}
