package sort

import "fmt"

// https://blog.csdn.net/yangh13/article/details/84564641/  父子位置关系证明：2i+1， 2i+2

// https://www.cnblogs.com/chengxiao/p/6129630.html 堆是完全二叉树
func heapSort(arr []int) {
	swap := func(arr []int, i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr[i:])
	}
	//fmt.Println(arr)
	for j := len(arr) - 1; j > 0; j-- {
		swap(arr, 0, j) // 不断将大顶堆的root交换到数组尾部，尾部的这些不再参加下一轮的堆调整
		//fmt.Println(arr)
		adjustHeap(arr[0:j])
	}
}

// 调整大顶堆（仅是因root节点变化而需要的调整过程，建立在大顶堆已构建的基础上）
func adjustHeap(arr []int) {
	i := 0
	for k := 2*i + 1; k < len(arr); k = k*2 + 1 {
		if k+1 < len(arr) && arr[k] < arr[k+1] {
			k++
		}
		if arr[k] > arr[i] {
			arr[k], arr[i] = arr[i], arr[k]
			i = k
		} else {
			break
		}
	}
}

func HeapSortTest() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heapSort(arr)
	fmt.Println(arr)
}
