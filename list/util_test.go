package list

import (
	"testing"
)

func Test_genList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	printList(genList(arr))
}
