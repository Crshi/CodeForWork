package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if A == nil || B == nil {
		return false
	}

	return dns(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func dns(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil {
		return false
	}

	if A.Val == B.Val {
		return dns(A.Left, B.Left) && dns(A.Right, B.Right)
	} else {
		return false
	}

	return false
}

func main() {

}
