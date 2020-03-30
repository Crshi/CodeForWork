package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	datas := make([]int, 0)
	dfs(root, &datas)
	return datas[k-1]
}

func dfs(root *TreeNode, datas *[]int) {
	if root.Right != nil {
		dfs(root.Right, datas)
	}

	*datas = append(*datas, root.Val)

	if root.Left != nil {
		dfs(root.Left, datas)
	}
}

func main() {

}
