package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	data := make([][]int, 0)

	if root == nil {
		return data
	}

	queue := []*TreeNode{root}
	tag := 0

	for len(queue) > 0 {
		length := len(queue)
		data = append(data, []int{})
		for i := 0; i < length; i++ {
			if tag&1 == 0 {
				data[tag] = append(data[tag], queue[i].Val)
			} else {
				var s = []int{queue[i].Val}
				data[tag] = append(s, data[tag]...)
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		tag++
		queue = queue[length:]
	}

	return data
}

func main() {

}
