package main

import "fmt"

type node struct {
	val  int
	prev *node
	next *node
}

type lruCache struct {
	head       *node
	tail       *node
	ht         map[int]*node
	numElement int
	capacity   int
}

func newLruCache(capacity int) *lruCache {
	return &lruCache{
		head:       nil,
		tail:       nil,
		ht:         make(map[int]*node),
		numElement: 0,
		capacity:   capacity,
	}
}

func (cache *lruCache) get(key int) int {
	node, ok := cache.ht[key]
	if !ok {
		return -1
	}
	cache.detach(node)
	cache.addToHead(node)
	return node.val
}

func (cache *lruCache) put(key, val int) {
	target, ok := cache.ht[key]
	if ok {
		target.val = val
		cache.addToHead(target)
		return
	}

	newNode := &node{val, nil, nil}
	cache.ht[key] = newNode
	if cache.head == nil && cache.tail == nil {
		cache.head = newNode
		cache.tail = newNode
		cache.numElement++
	} else if cache.numElement+1 > cache.capacity {
		delete(cache.ht, cache.tail.val)
		if cache.tail.prev == nil {
			cache.head = newNode
			cache.tail = newNode
		} else {
			cache.tail.prev.next = nil
			cache.tail = cache.tail.prev
			cache.addToHead(newNode)
		}
	} else {
		cache.addToHead(newNode)
		cache.numElement++
	}
}

func (cache *lruCache) detach(node *node) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if cache.head == node {
		cache.head = node.next
	} else if cache.tail == node {
		cache.tail = node.prev
	}
}

func (cache *lruCache) addToHead(node *node) {
	node.next = cache.head
	if cache.head != nil {
		cache.head.prev = node
	}
	cache.head = node
}

func main() {
	lru := newLruCache(4)
	lru.put(1, 1)
	//fmt.Println(lru.head.val)
	//fmt.Println(lru.tail.val)
	lru.put(2, 2)
	//fmt.Println(lru.head.val)
	//fmt.Println(lru.tail.val)
	fmt.Println(lru.get(1))
	//fmt.Println(lru.head.val)
	//fmt.Println(lru.tail.val)
	lru.put(3, 3)
	fmt.Println(lru.get(2))
	lru.put(4, 4)
	fmt.Println(lru.get(1))
	fmt.Println(lru.get(3))
	fmt.Println(lru.get(4))
}
