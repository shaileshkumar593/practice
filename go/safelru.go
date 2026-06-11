package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key   int
	value int

	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int

	cache map[int]*Node

	head *Node
	tail *Node

	mu sync.Mutex
}

func Constructor(capacity int) *LRUCache {

	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) remove(node *Node) {

	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) insertFront(node *Node) {

	node.next = l.head.next
	node.prev = l.head

	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) moveToFront(node *Node) {
	l.remove(node)
	l.insertFront(node)
}

func (l *LRUCache) Get(key int) int {

	l.mu.Lock()
	defer l.mu.Unlock()

	node, ok := l.cache[key]
	if !ok {
		return -1
	}

	l.moveToFront(node)

	return node.value
}

func (l *LRUCache) Put(key, value int) {

	l.mu.Lock()
	defer l.mu.Unlock()

	// update existing
	if node, ok := l.cache[key]; ok {

		node.value = value
		l.moveToFront(node)
		return
	}

	// eviction
	if len(l.cache) == l.capacity {

		lru := l.tail.prev

		delete(l.cache, lru.key)
		l.remove(lru)
	}

	node := &Node{
		key:   key,
		value: value,
	}

	l.insertFront(node)
	l.cache[key] = node
}

func main() {

	cache := Constructor(2)

	cache.Put(1, 10)
	cache.Put(2, 20)

	fmt.Println(cache.Get(1))

	cache.Put(3, 30)

	fmt.Println(cache.Get(2))
}