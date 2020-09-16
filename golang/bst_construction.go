package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type binarySearchTreeOperation interface {
	insert(value int) *BST
	contains(value int) bool
	remove(value int) *BST
}

func (tree *BST) insert(value int) *BST {
	// case 1. the tree is nil, insert it into that node
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.insert(value)
		}
	}
	return tree
}

func (tree *BST) contains(value int) bool {
	if value < tree.Value {
		if tree.Left == nil {
			return false
		} else {
			return tree.Left.contains(value)
		}
	} else if value > tree.Value {
		if tree.Right == nil {
			return false
		} else {
			return tree.Right.contains(value)
		}
	}
	return true
}

func (tree *BST) remove(value int) *BST {
	tree.removeHelper(value, nil)
	return tree
}

func (tree *BST) removeHelper(value int, parent *BST) {
	current := tree
	for current != nil {
		if value < current.Value {
			parent = current 
			current = current.Left
		} else if value > current.Value {
			parent = current
			current = current.Right
		} else {
			if current.Left != nil && current.Right != nil {
				current.Value = current.Right.getMinValue()
				current.Right.removeHelper(current.Value, current)
			} else if parent == nil {
				if current.Left != nil {
					current.Value = current.Left.Value
				}
			}
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}