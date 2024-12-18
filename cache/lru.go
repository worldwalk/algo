package cache

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List // 标准库的双向链表，值得学习其源码
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*Pair).value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value.(*Pair).value = value
	} else {
		if l.list.Len() == l.capacity {
			tail := l.list.Back()
			delete(l.cache, tail.Value.(*Pair).key)
			l.list.Remove(tail)
		}
		elem := l.list.PushFront(&Pair{key, value})
		l.cache[key] = elem
	}
}

func LRUTest() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // Output: -1
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
