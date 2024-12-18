package list

func ReverseTest() {
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: nil}}}}}
	printList(head)
	//printList(reverse(head))
	//printList(reverse2(head))
	printList(reverse3(head))
}
func reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverse2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &Node{Next: head}
	for next := head.Next; next != nil; {
		curr := next
		next = next.Next
		curr.Next = dummy.Next
		dummy.Next = curr
	}
	head.Next = nil
	return dummy.Next
}

func reverse3(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	var pre *Node = nil
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
