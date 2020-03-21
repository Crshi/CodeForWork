package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	tmp := head

	relation := make(map[*Node]*Node, 0)
	for head != nil {
		tmp := Node{
			Val: head.Val,
		}
		relation[head] = &tmp
		head = head.Next
	}

	t := tmp

	for t != nil {
		if relation[t.Next] != nil {
			relation[t].Next = relation[t.Next]
		}

		if relation[t.Random] != nil {
			relation[t].Random = relation[t.Random]
		}

		t = t.Next
	}

	return relation[tmp]
}

func main() {

	node3 := Node{
		Val: 3,
	}

	node2 := Node{
		Val:  2,
		Next: &node3,
	}

	node1 := Node{
		Val:  1,
		Next: &node2,
	}

	node2.Random = &node1

	t := copyRandomList(&node1)

	for t != nil {
		println(t.Val)
		t = t.Next
	}
}
