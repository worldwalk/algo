package list

import "fmt"

func printList(head *Node) {
	for ; head != nil; head = head.Next {
		fmt.Print(head.Val)
	}
	fmt.Println()
}

func genList(arr []int) *Node {
	dummy := &Node{}
	p := dummy
	for _, v := range arr {
		node := &Node{Val: v}
		p.Next = node
		p = p.Next
	}
	return dummy.Next
}

func getLen(head *Node) int {
	ret := 0
	for head != nil {
		head = head.Next
		ret++
	}
	return ret
}
