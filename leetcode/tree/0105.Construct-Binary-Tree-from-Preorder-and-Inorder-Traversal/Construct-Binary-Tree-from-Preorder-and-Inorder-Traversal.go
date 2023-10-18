package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	mid := 0
	for i, val := range inorder {
		if val == root.Val {
			mid = i
			break
		}
	}
	//lftInorder := inorder[:mid]
	//rgtInorder := inorder[mid+1:]
	//lftPreorder := preorder[1 : 1+mid]
	//rgtPreorder := preorder[1+mid:]
	root.Left = buildTree(preorder[1:1+mid], inorder[:mid])
	root.Right = buildTree(preorder[1+mid:], inorder[1+mid:])

	return root
}

func main() {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	root := buildTree(preOrder, inOrder)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)
	fmt.Println(root.Right.Left.Val)
	fmt.Println(root.Right.Right.Val)
}
