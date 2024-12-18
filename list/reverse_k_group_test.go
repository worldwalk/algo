package list

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: nil}}}}}
	printList(head)
	printList(reverseKGroup(head, 2))
}

func Test_reReverseKGroup(t *testing.T) {
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: nil}}}}}
	printList(head)
	printList(reReverseKGroup(head, 2))
}
