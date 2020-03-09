package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}
	index := 0

	for i, x := range inorder {
		if x == preorder[0] {
			index = i
			break
		}
	}

	node.Left = buildTree(preorder[1:index+1], inorder[:index])
	node.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return node
}

//Preorder 的Index位置是怎么来的
//Index 是当前根节点左子树的节点数 谢谢

func main() {
	println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}).Val)
}
