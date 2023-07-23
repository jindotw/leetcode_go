package schema

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) GetLeft() *TreeNode {
	return node.Left
}

func (node *TreeNode) GetRight() *TreeNode {
	return node.Right
}

func (node *TreeNode) GetVal() int {
	return node.Val
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

func MakeTreePathSum() *TreeNode {
	node41 := &TreeNode{Val: 3}
	node42 := &TreeNode{Val: -2}
	node43 := &TreeNode{Val: 1}

	node31 := &TreeNode{3, node41, node42}
	node32 := &TreeNode{2, nil, node43}
	node33 := &TreeNode{11, nil, nil}

	node21 := &TreeNode{5, node31, node32}
	node22 := &TreeNode{-3, nil, node33}

	return &TreeNode{10, node21, node22}
}

func PreOrderRecursion(node *TreeNode, val *[]int) {
	if node == nil {
		return
	}

	*val = append(*val, node.GetVal())
	PreOrderRecursion(node.GetLeft(), val)
	PreOrderRecursion(node.GetRight(), val)
}
