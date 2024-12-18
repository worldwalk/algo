package list

// reReverseKGroup 从后往前K个一组翻转链表
func reReverseKGroup(head *Node, k int) *Node {
	if k <= 1 || head == nil {
		return head
	}
	length := getLen(head)
	if length < k {
		return head
	}

	if length%k == 0 {
		return reverseKGroup(head, k)
	}

	cur := head
	pre := head
	for i := 0; i < length%k; i++ {
		pre = cur
		cur = cur.Next
	}
	pre.Next = reverseKGroup(cur, k)
	return head
}

// reverseKGroup K个一组翻转链表
func reverseKGroup(head *Node, k int) *Node {
	if k <= 1 || head == nil {
		return head
	}
	i := 0
	var l, r *Node
	l = head
	cur := head
	for ; i < k && cur != nil; i++ {
		cur = cur.Next
	}
	if i < k {
		return head
	}
	r = cur
	newHead := reverseBtw(l, r)
	l.Next = reverseKGroup(r, k)
	return newHead
}

// [head, tail)
func reverseBtw(head, tail *Node) *Node {
	dummy := &Node{}
	for cur := head; cur != tail; {
		next := cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
		cur = next
	}
	return dummy.Next
}
