package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	// key is the original node, value is the cloned version
	set := make(map[*Node]*Node)
	var dfs func(*Node) *Node
	dfs = func(vertex *Node) *Node {
		if cached, ok := set[vertex]; ok {
			return cached
		}
		retNode := &Node{Val: vertex.Val}
		set[vertex] = retNode
		for _, neighbour := range vertex.Neighbors {
			retNode.Neighbors = append(retNode.Neighbors, dfs(neighbour))
		}

		return retNode
	}

	return dfs(node)
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)
	cloned := cloneGraph(node1)
	fmt.Println(cloned)
}
