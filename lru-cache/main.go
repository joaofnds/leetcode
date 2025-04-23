package main

import "container/list"

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element, capacity),
	}
}

func (lru *LRUCache) Get(key int) int {
	val, ok := lru.cache[key]
	if !ok {
		return -1
	}

	lru.list.MoveToFront(val)
	return val.Value.([]int)[1]
}

func (lru *LRUCache) Put(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		lru.list.Remove(elem)
		lru.cache[key] = lru.list.PushFront([]int{key, value})
		return
	}

	if len(lru.cache) == lru.capacity {
		toRemove := lru.list.Remove(lru.list.Back())
		delete(lru.cache, toRemove.([]int)[0])
	}

	lru.cache[key] = lru.list.PushFront([]int{key, value})
}
