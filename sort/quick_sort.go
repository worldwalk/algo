package sort

import "fmt"

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition3(arr, left, right)
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func partition2(arr []int, left, right int) int { // 这个版本好理解，i代表下一个available slot
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// 二路快排 版本的partition
func partition3(arr []int, left, right int) int {
	i, j := left, right
	pivot := arr[i]
	for i < j {
		// 必须先移动j指针，否则不会正确排序，why？
		// 因为：是为了保证比pivot大的数必须在pivot右边！ 若先走i，走完后arr[i]> pivot，然后交换了，导致最终arr[left]> pivot
		// 这个例子走一轮就明白了： arr := []int{5, 3, 8, 4, 2, 7, 1, 6, 9}
		for i < j && arr[j] >= pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left], arr[i] = arr[i], arr[left]
	return i
}

func QuickSortTest() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6, 9}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr) // 输出：[1 2 3 4 5 6 7 8]

}
