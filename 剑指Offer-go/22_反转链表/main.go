package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	head.Next = nil

	for tmp != nil {
		if tmp.Next == nil {
			tmp.Next = head
			return tmp
		} else {
			t := tmp.Next
			tmp.Next = head
			head = tmp
			tmp = t
		}
	}

	return head
}

func main() {

}
