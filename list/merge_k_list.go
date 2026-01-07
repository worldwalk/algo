package list

import (
	"container/heap"
	"math"
)

/**
朴素迭代法：时间复杂度稳定\(O(Nk)\), 仅适用于 k 极小（如 k≤5）的场景，优点是代码极简、零额外空间，缺点是性能随 k 增大急剧下降；
归并法：递归版空间\(O(logk)\)、迭代版空间\(O(1)\)，时间复杂度稳定\(O(N logk)\)，是 k 较大时的高效选择，代码可读性也较好；
堆优化迭代法：工业界首选，时间\(O(N logk)\)、空间\(O(k)\)，支持动态链表合并，稳定性优于递归归并法；
性能拐点：当 k≥10 时，归并法 / 堆优化法的性能就已显著优于朴素迭代法。
*/
// 合并k个有序链表 - 迭代
func mergeKListsByIter(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	dummy := &Node{}
	cur := dummy
	for {
		idx := -1
		minVal := math.MaxInt32
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && lists[i].Val < minVal {
				idx = i
				minVal = lists[i].Val
			}
		}
		if idx == -1 {
			break
		}
		cur.Next = lists[idx]
		cur = cur.Next
		lists[idx] = lists[idx].Next
	}
	cur.Next = nil
	return dummy.Next
}

// 合并k个有序链表 - 归并
func mergeKLists(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)
}

// 合并两个有序链表
func mergeTwoLists(l1, l2 *Node) *Node {
	dummy := &Node{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

// 堆优化的迭代法（工业界常用）
func mergeKListsByHeap(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}
	// 初始化最小堆
	h := &NodeHeap{}
	heap.Init(h)
	// 把每个链表的头节点加入堆
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}
	// 合并逻辑
	dummy := &Node{}
	cur := dummy
	for h.Len() > 0 {
		// 取出堆顶（最小值节点）
		minNode := heap.Pop(h).(*Node)
		cur.Next = minNode
		cur = cur.Next
		// 该节点的下一个节点入堆
		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}
	return dummy.Next
}

// 定义优先队列（最小堆）
type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
