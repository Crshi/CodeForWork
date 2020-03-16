package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1

	if head == nil {
		head = l2
	}

	if l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			head = l2
		} else {
			head = l1
		}
	}

	if head == nil {
		return head
	}

	if head == l1 {
		l1 = l1.Next
	} else {
		l2 = l2.Next
	}

	tmp := head

	for {
		if l1 == nil {
			if l2 == nil {
				return head
			} else {
				tmp.Next = l2
				tmp = tmp.Next
				l2 = l2.Next
			}
		} else {
			if l2 == nil {
				tmp.Next = l1
				tmp = tmp.Next
				l1 = l1.Next
			} else {
				if l1.Val > l2.Val {
					tmp.Next = l2
					tmp = tmp.Next
					l2 = l2.Next
				} else {
					tmp.Next = l1
					tmp = tmp.Next
					l1 = l1.Next
				}
			}
		}
	}

	return head
}

func main() {

}
