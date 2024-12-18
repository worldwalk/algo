package sort

import "fmt"

func MergeSortTest() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6}
	arr = mergeSort(arr)
	fmt.Println(arr) // 输出：[1 2 3 4 5 6 7 8]
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var ret []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}

	ret = append(ret, left[i:]...)
	ret = append(ret, right[j:]...)
	return ret
}
