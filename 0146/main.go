package main

/*
	KEY IDEA:
	Use a doubly-linkedlist to determine most infrequently used node.
	Nodes should hold the key and value.

	Map should point to nodes in the LinkedList to allow constant lookup.
	Move to front if used.
*/

type node struct {
	next  *node
	prev  *node
	key   int
	value int
}

type linkedList struct {
	head *node
	tail *node
}

func (this *linkedList) push(key, value int) *node {

	newNode := node{this.head, nil, key, value}

	if this.head != nil {
		this.head.prev = &newNode
	}
	if this.tail == nil {
		this.tail = &newNode
	}

	this.head = &newNode

	return &newNode
}

func (this *linkedList) pop() (int, int) {
	node := this.tail

	k := node.key
	v := node.value

	if this.tail.prev == nil {
		this.head = nil
		this.tail = nil
	} else {
		this.tail.prev.next = nil
		this.tail = this.tail.prev
	}

	return k, v
}

func (this *linkedList) moveToFront(node *node) {
	if node.prev != nil {
		node.prev.next = node.next
		if node.next != nil {
			node.next.prev = node.prev
		} else {
			this.tail = node.prev
		}
		node.next = this.head
		node.prev = nil
		this.head.prev = node
		this.head = node
	}
}

type LRUCache struct {
	cache       map[int]*node
	list        linkedList
	capacity    int
	curCapacity int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{make(map[int]*node), linkedList{nil, nil}, capacity, 0}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.cache[key]; ok {
		this.list.moveToFront(val)
		return val.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.cache[key]; ok {
		val.value = value
		this.list.moveToFront(val)
	} else {
		if this.curCapacity >= this.capacity {
			k, _ := this.list.pop()
			delete(this.cache, k)
		} else {
			this.curCapacity++
		}
		node := this.list.push(key, value)
		this.cache[key] = node
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Get(3)
	cache.Get(4)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
