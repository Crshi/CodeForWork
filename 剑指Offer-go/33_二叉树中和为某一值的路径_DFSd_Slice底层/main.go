package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	dfsSearchTree(root, sum, []int{}, &res)

	return res
}

func dfsSearchTree(root *TreeNode, sum int, tmp []int, res *[][]int) {
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil {
		t := 0
		for _, x := range tmp {
			t += x
		}

		if t == sum {
			//slice是一个指向底层的数组的指针结构体
			//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
			//所以这里需要 copy arr 到 tmp，再添加进 ret，防止 arr 底层数据修改带来的错误
			t := make([]int, len(tmp))
			copy(t, tmp)
			*res = append(*res, t)
		}

		return
	} else {
		if root.Left != nil {
			dfsSearchTree(root.Left, sum, tmp, res)
		}

		if root.Right != nil {
			dfsSearchTree(root.Right, sum, tmp, res)
		}
	}
}

func main() {
	head := TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 8,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
				Left: &TreeNode{
					Val: 5,
				},
			},
			Left: &TreeNode{
				Val: 13,
			},
		},
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	}

	data := pathSum(&head, 22)

	for _, x := range data {
		fmt.Println(x)
	}
}
