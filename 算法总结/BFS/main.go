package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	datas := make([]int, 0)
	for len(queue) > 0 {
		tmp := len(queue)
		for i := 0; i < tmp; i++ {
			datas = append(datas, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[tmp:]
	}

	return datas
}

func main() {

}
