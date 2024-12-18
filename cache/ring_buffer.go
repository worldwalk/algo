package cache

import "fmt"

type RingBuffer struct {
	buffer []int
	size   int
	head   int
	tail   int
	count  int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, size),
		size:   size,
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (rb *RingBuffer) Enqueue(value int) {
	if rb.count == rb.size {
		fmt.Println("Ring buffer is full")
		return
	}

	rb.buffer[rb.tail] = value
	rb.tail = (rb.tail + 1) % rb.size
	rb.count++
}

func (rb *RingBuffer) Dequeue() int {
	if rb.count == 0 {
		fmt.Println("Ring buffer is empty")
		return -1
	}

	value := rb.buffer[rb.head]
	rb.head = (rb.head + 1) % rb.size
	rb.count--
	return value
}
