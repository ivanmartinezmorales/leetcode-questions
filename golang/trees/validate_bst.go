package main

import "math"

type BST struct {
		Value int
		Left *BST
		Right *BST
}

func (tree *BST) Validate() bool {
		return tree.validate(math.MinInt32, math.MaxInt32);
}

func (tree *BST) validate(min, max int) bool {
		if tree == nil {
				return true	
		}

		if tree.Value < min || tree.Value >= max {
				return false
		}
		
		return tree.Left.validate(min, tree.Value) && tree.Right.validate(tree.Value, max)
}