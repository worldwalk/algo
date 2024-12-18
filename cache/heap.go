package cache

import "container/heap"

// see example in same dir as heap.Interface, 可以实现priorityQueue，见同目录下的example_pq_test.go

type heapImpl struct {
}

func NewHeap() heap.Interface {
	return &heapImpl{}
}
func (h heapImpl) Len() int {
	//TODO implement me
	panic("implement me")
}

func (h heapImpl) Less(i, j int) bool {
	//TODO implement me
	panic("implement me")
}

func (h heapImpl) Swap(i, j int) {
	//TODO implement me
	panic("implement me")
}

func (h heapImpl) Push(x any) {
	//TODO implement me
	panic("implement me")
}

func (h heapImpl) Pop() any {
	//TODO implement me
	panic("implement me")
}
