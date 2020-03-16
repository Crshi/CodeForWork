package main

import (
	"fmt"
)

// 题目类型：树的遍历，递归，构造
// 先序遍历给出的二叉搜索树，按照中序遍历的顺序将树中元素保存下来。
// 二叉搜索树的先序遍历是升序的(从小到大)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	data := make([]int, 0)
	dfs(root, &data)
	return construct(data, 0, len(data)-1)
}

// 通过DFS根据升序数组构造平衡二叉搜索树保存到data中：
func dfs(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}

	//先遍历左子树
	dfs(root.Left, data)
	//写入数组
	*data = append(*data, root.Val)
	//后遍历右子树保证顺序
	dfs(root.Right, data)
}

// 如果数组为空，则对应的树亦为空。
// 如果数组不为空，设长度为 n，那么位置n/2处的元素应为树的根节点。子数组 [1, n/2-1] 及 [n/2+1, n]分别对应左右子树。
// 因为两个子数组的长度相差不会超过 1，所以保证了左右子树的高度相差不会超过 1。
func construct(data []int, L int, R int) *TreeNode {
	if L > R {
		return nil
	}

	mid := (L + R) >> 1
	ptr := TreeNode{
		Val:   data[mid],
		Right: construct(data, mid+1, R),
		Left:  construct(data, L, mid-1),
	}

	return &ptr
}

func initTree() *TreeNode {

	head4 := TreeNode{
		Val: 4,
	}

	head3 := TreeNode{
		Val:   3,
		Right: &head4,
	}

	head2 := TreeNode{
		Val:   2,
		Right: &head3,
	}

	head := TreeNode{
		Val:   1,
		Right: &head2,
	}

	return &head
}

func main() {
	head := initTree()

	fmt.Println(balanceBST(head))
}
