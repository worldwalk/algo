package sort

import "fmt"

func BubbleSortTest() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6}
	//BubbleSort(arr)
	BubbleSortPro(arr)
	fmt.Println(arr) // 输出：[1 2 3 4 5 6 7 8]
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSortPro(arr []int) {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !flag {
			break
		}
	}
}
