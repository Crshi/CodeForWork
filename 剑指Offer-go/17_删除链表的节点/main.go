package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {

	if head.Val == val {
		return head.Next
	}

	tmp := head
	last := head

	for head.Next != nil {
		if tmp.Val == val {
			last.Next = tmp.Next
			break
		}
		last = tmp
		tmp = tmp.Next
	}

	return head
}

func initListNode(head *ListNode) {
	for i := 1; i <= 3; i++ {
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

	data := deleteNode(&head, 1)

	for data != nil {
		println(data.Val)
		data = data.Next
	}
}
