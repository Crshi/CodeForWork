package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res, _ := helper(root)
	return res
}

func helper(root *TreeNode) (res bool, deepth int) {
	if root == nil {
		return true, 0
	}

	resleft, left := helper(root.Left)
	resright, right := helper(root.Right)

	if resleft == false || resright == false || abs(left, right) > 1 {
		return false, 0
	}

	return true, max(left, right) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func main() {

}
