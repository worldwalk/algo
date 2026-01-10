package list

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	// 测试用例1: 多个有序链表
	lists := []*Node{
		genList([]int{1, 4, 5}),
		genList([]int{1, 3, 4}),
		genList([]int{2, 6}),
	}
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}
	result := mergeKLists(lists)
	if !equal(result, expected) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected, toSlice(result))
	}

	// 测试用例2: 空链表
	lists = []*Node{}
	result = mergeKLists(lists)
	if result != nil {
		t.Errorf("Test case 2 failed: expected nil, got %v", toSlice(result))
	}

	// 测试用例3: 单个链表
	lists = []*Node{genList([]int{1, 2, 3})}
	expected = []int{1, 2, 3}
	result = mergeKLists(lists)
	if !equal(result, expected) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected, toSlice(result))
	}

	// 测试用例4: 包含空链表
	lists = []*Node{
		nil,
		genList([]int{1, 2, 3}),
		nil,
		genList([]int{4, 5, 6}),
	}
	expected = []int{1, 2, 3, 4, 5, 6}
	result = mergeKLists(lists)
	if !equal(result, expected) {
		t.Errorf("Test case 4 failed: expected %v, got %v", expected, toSlice(result))
	}
}

func Test_mergeKListsByHeap(t *testing.T) {
	// 测试用例1: 多个有序链表
	lists := []*Node{
		genList([]int{1, 4, 5}),
		genList([]int{1, 3, 4}),
		genList([]int{2, 6}),
	}
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}
	result := mergeKListsByHeap(lists)
	if !equal(result, expected) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected, toSlice(result))
	}

	// 测试用例2: 空链表
	lists = []*Node{}
	result = mergeKListsByHeap(lists)
	if result != nil {
		t.Errorf("Test case 2 failed: expected nil, got %v", toSlice(result))
	}

	// 测试用例3: 单个链表
	lists = []*Node{genList([]int{1, 2, 3})}
	expected = []int{1, 2, 3}
	result = mergeKListsByHeap(lists)
	if !equal(result, expected) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected, toSlice(result))
	}

	// 测试用例4: 包含空链表
	lists = []*Node{
		nil,
		genList([]int{1, 2, 3}),
		nil,
		genList([]int{4, 5, 6}),
	}
	expected = []int{1, 2, 3, 4, 5, 6}
	result = mergeKListsByHeap(lists)
	if !equal(result, expected) {
		t.Errorf("Test case 4 failed: expected %v, got %v", expected, toSlice(result))
	}
}

// 辅助函数：将链表转换为切片
func toSlice(head *Node) []int {
	var slice []int
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}

// 辅助函数：比较链表和切片是否相等
func equal(head *Node, expected []int) bool {
	slice := toSlice(head)
	if len(slice) != len(expected) {
		return false
	}
	for i := range slice {
		if slice[i] != expected[i] {
			return false
		}
	}
	return true
}
