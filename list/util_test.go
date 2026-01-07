package list

import (
	"testing"
)

func Test_genList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	printList(genList(arr))
}

func Test_isMinLen(t *testing.T) {
	tests := []struct {
		name string
		head *Node
		k    int
		want bool
	}{{
		name: "list length equals k",
		head: genList([]int{1, 2, 3}),
		k:    3,
		want: true,
	}, {
		name: "list length greater than k",
		head: genList([]int{1, 2, 3, 4, 5}),
		k:    3,
		want: true,
	}, {
		name: "list length less than k",
		head: genList([]int{1, 2}),
		k:    3,
		want: false,
	}, {
		name: "empty list",
		head: nil,
		k:    1,
		want: false,
	}, {
		name: "k equals 0",
		head: genList([]int{1, 2, 3}),
		k:    0,
		want: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMinLen(tt.head, tt.k); got != tt.want {
				t.Errorf("isMinLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
