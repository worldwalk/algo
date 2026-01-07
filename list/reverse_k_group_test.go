package list

import (
	"testing"
)

// Test_reverseKGroup 测试K个一组翻转链表（从前往后）
func Test_reverseKGroup(t *testing.T) {
	// 创建测试链表: 1->2->3->4->5
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: nil}}}}}

	// 打印测试信息
	t.Log("=== Test_reverseKGroup ===")
	t.Log("输入链表: 1->2->3->4->5")
	t.Log("K=2分组翻转")
	t.Log("原始链表:")
	printList(head)

	// 执行翻转
	result := reverseKGroup(head, 2)

	// 打印结果
	t.Log("翻转后链表:")
	printList(result)
	t.Log("预期结果: 2->1->4->3->5")
}

// Test_reReverseKGroup 测试K个一组翻转链表（从后往前）
func Test_reReverseKGroup(t *testing.T) {
	// 创建测试链表: 1->2->3->4->5
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: nil}}}}}

	// 打印测试信息
	t.Log("=== Test_reReverseKGroup ===")
	t.Log("输入链表: 1->2->3->4->5")
	t.Log("从后往前K=2分组翻转")
	t.Log("原始链表:")
	printList(head)

	// 执行翻转
	result := reReverseKGroup(head, 2)

	// 打印结果
	t.Log("翻转后链表:")
	printList(result)
	t.Log("预期结果: 1->3->2->5->4")
}
