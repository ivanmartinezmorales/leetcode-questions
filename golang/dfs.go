package main

func (n *Node) depthFirstSearch(arr []string) []string {
	arr = append(arr, n.Name)
	for _, child := range n.Children {
		arr = child.depthFirstSearch(arr)
	}
	return arr
}
