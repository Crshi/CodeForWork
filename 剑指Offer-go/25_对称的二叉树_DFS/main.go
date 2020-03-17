package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isChildSymmetric(root.Left, root.Right)
}

func isChildSymmetric(Left *TreeNode, Right *TreeNode) bool {
	if Left == nil && Right == nil {
		return true
	}

	if Left == nil || Right == nil || Left.Val != Right.Val {
		return false
	}

	return isChildSymmetric(Left.Left, Right.Right) && isChildSymmetric(Left.Right, Right.Left)
}

func main() {

}
