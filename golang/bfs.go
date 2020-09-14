package main

type Node struct {
		Name string
		Children []*Node
}

func (n *Node) BFS(nodes []string) []string {
		queue := []*Node{n}
		for len(queue) > 0 {
				// Popping the node off the queue
				current := queue[0]
				queue = queue[1:]
				nodes = append(nodes, current.Name)
				for _, child := range current.Children {
						queue = append(queue, child)
				}
		}
		return nodes
}
