package trees

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	tree.Left, tree.Right = tree.Right, tree.Left // swapping nodes
	if tree.Left != nil {
		tree.Left.InvertBinaryTree()
	}

	if tree.Right != nil {
		tree.Right.InvertBinaryTree()
	}

}
