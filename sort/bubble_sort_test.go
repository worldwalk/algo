package sort

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSortPro(b *testing.B) {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSortPro(arr)
	}
}
