package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	max := &m //现有的最长直径
	if root == nil {
		return 0
	}
	depth(root, max)
	return *max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	a := depth(root.Left, max)
	b := depth(root.Right, max)
	*max = int(math.Max(float64(a+b), float64(*max))) //此节点的直径和现有的最长直径比较
	return int(math.Max(float64(a), float64(b))) + 1  //返回深度
}

func main() {

	node5 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node4 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   2,
		Left:  &node4,
		Right: &node5,
	}

	head := TreeNode{
		Val:   1,
		Left:  &node2,
		Right: &node3,
	}

	println(diameterOfBinaryTree(&head))
}
