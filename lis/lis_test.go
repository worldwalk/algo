package lis

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {

	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "normal1", arr: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4}, //[2,3,7,101]
		{name: "normal2", arr: []int{0, 1, 0, 3, 2, 3}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.arr); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printLISPath(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{name: "normal1", arr: []int{10, 9, 2, 5, 3, 7, 101, 18}}, //[2,3,7,101]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printLISPath(tt.arr)
		})
	}
}
