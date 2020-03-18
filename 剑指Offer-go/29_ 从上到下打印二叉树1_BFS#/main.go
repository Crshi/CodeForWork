package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var nums []int

	for len(queue) > 0 {
		t := queue[0]
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
		nums = append(nums, t.Val)
		queue = queue[1:]
	}
	return nums
}

func main() {

}
