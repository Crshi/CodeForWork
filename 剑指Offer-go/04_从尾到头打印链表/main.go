package main

type ListNode struct {
	 Val int
	 Next *ListNode
 }

func reversePrint(head *ListNode) []int {

	datas := []int{}

	for head != nil {
		datas = append(datas,head.Val)
		head = head.Next
	}

	data := []int{}

	for i := len(datas) -1;i >= 0 ;i-- {
		data = append(data,datas[i])
	}

	return data
}

func initListNode(head *ListNode){
	for i := 1;i <= 3;i++{
		next := ListNode{
			Val:i,
			Next:nil,
		}

		head.Next = &next
		head = &next
	}
}

func main() {

	head := ListNode {
		Val:0,
		Next:nil,
	}

	initListNode(&head)

	datas := reversePrint(&head)
	println(datas)
}
