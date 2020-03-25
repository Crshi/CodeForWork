package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Encodes a tree to a single string.
func serialize(root *TreeNode) string {
	res := "["
	if root == nil {
		return "[]"
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i] == nil {
				res = res + "null,"
				continue
			}
			res = res + strconv.Itoa(queue[i].val) + ","
			queue = append(queue, queue[i].left)
			queue = append(queue, queue[i].right)
		}

		queue = queue[length:]
	}

	for {
		if res[len(res)-5:] == "null," {
			res = res[:len(res)-5]
		} else {
			break
		}
	}

	res = res[:len(res)-1] + "]"

	return res
}

//Decodes your encoded data to tree.
func deserialize(data string) *TreeNode {
	data = data[1 : len(data)-1]
	datas := strings.Split(data, ",")

	head := TreeNode{
		val: strconv.Atoi(datas[0]),
	}

	for i := 1; i < len(datas); i++ {

	}
}

func main() {
	head := TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,
		},
		right: &TreeNode{
			val: 3,
			left: &TreeNode{
				val: 4,
			},
			right: &TreeNode{
				val: 5,
			},
		},
	}

	head2 := deserialize(serialize(&head))

	println(head2)
}
