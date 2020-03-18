package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var nums [][]int

	for len(queue) > 0 {
		var tmp []int
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		nums = append(nums, tmp)
	}
	return nums
}

func main() {

}
