package dyopro

import "testing"

func Test_knapsack01(t *testing.T) {
	tests := []struct {
		name     string
		w        []int
		val      []int
		capacity int
		want     int
	}{
		{name: "normal1", w: []int{2, 1, 3}, val: []int{4, 2, 3}, capacity: 5, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack01(tt.w, tt.val, tt.capacity); got != tt.want {
				t.Errorf("knapsack01() = %v, want %v", got, tt.want)
			}
		})
	}
}
