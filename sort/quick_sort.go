package sort

import "fmt"

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition2(arr, left, right)
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

func partition2(arr []int, left, right int) int {
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

func QuickSortTest() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6, 9}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr) // 输出：[1 2 3 4 5 6 7 8]

}
