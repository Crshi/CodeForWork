package main

 type ListNode struct {
	 Val int
	 Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	tmp := head
	nodes := []*ListNode{head}

	for {
		if(tmp.Next != nil) {
			tmp = tmp.Next
			nodes = append(nodes,tmp)
		} else {
			break
		}
	}

	return nodes[len(nodes) - k]
}

func initListNode(head *ListNode) {
	for i := 1; i <= 5; i++ {
		next := ListNode{
			Val:  i,
			Next: nil,
		}

		head.Next = &next
		head = &next
	}
}

func main() {
	head := ListNode{
		Val:  0,
		Next: nil,
	}

	initListNode(&head)
	getKthFromEnd(&head,1)
}
