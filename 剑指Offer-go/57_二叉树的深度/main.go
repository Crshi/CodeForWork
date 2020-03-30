package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return getDepth(root, 0)
}

func getDepth(root *TreeNode, deepth int) int {
	if root == nil {
		return deepth
	} else {
		deepth++
		return max(getDepth(root.Left, deepth), getDepth(root.Right, deepth))
	}
}

func max(l int, r int) int {
	if l > r {
		return l
	}

	return r
}
