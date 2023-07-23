package main

import (
	"fmt"
	"leetcode_go/schema"
)

func pre() {
	root := schema.MakeTree()
	preOrderRecur(root)
	fmt.Println()
	preOrderIter(root)
	fmt.Println()
	preOrderCommon(root)
}

func in() {
	root := schema.MakeTree()
	inOrderRecur(root)
	fmt.Println()
	inOrderIter(root)
	fmt.Println()
	inOrderCommon(root)
	fmt.Println()
}

func post() {
	root := schema.MakeTree()
	postOrderRecur(root)
	fmt.Println()
	PostOrderCommon(root)
	fmt.Println()
	PostOrderIter(root)
}

func level() {
	root := schema.MakeTree()
	bfs(root)
}

func main() {
	level()
}
