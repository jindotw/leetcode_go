package schema

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree() *TreeNode {
	node33 := &TreeNode{Val: 3}
	node13 := &TreeNode{Val: 1}
	node53 := &TreeNode{Val: 5}

	node12 := &TreeNode{Val: 1, Left: node33}
	node42 := &TreeNode{Val: 4, Left: node13, Right: node53}

	node31 := &TreeNode{Val: 3, Left: node12, Right: node42}

	return node31
}
