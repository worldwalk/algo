package cache

import (
	"fmt"
	"testing"
)

func Test_ring_buffer(t *testing.T) {
	rb := NewRingBuffer(5)

	rb.Enqueue(1)
	rb.Enqueue(2)
	rb.Enqueue(3)
	rb.Enqueue(4)
	rb.Enqueue(5)

	fmt.Println(rb.Dequeue()) // Output: 1
	fmt.Println(rb.Dequeue()) // Output: 2

	rb.Enqueue(6)
	rb.Enqueue(7)

	fmt.Println(rb.Dequeue()) // Output: 3
	fmt.Println(rb.Dequeue()) // Output: 4
	fmt.Println(rb.Dequeue()) // Output: 5
	fmt.Println(rb.Dequeue()) // Output: 6
	fmt.Println(rb.Dequeue()) // Output: 7
	fmt.Println(rb.Dequeue()) // Output: Ring buffer is empty
}
