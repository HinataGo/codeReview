package lru

// 实现本题的两种操作，需要用到一个哈希表和一个双向链表。
// 在面试中，面试官一般会期望读者能够自己实现一个简单的双向链表，
// 而不是使用语言自带的、封装好的数据结构。

// 双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，
// 这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。

// 双向链表 hash 自己实现
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*dListNode
	head, tail *dListNode
}

// 手写双向链表
type dListNode struct {
	key, value int
	prev, next *dListNode
}

func initDList(k, v int) *dListNode {
	return &dListNode{
		key:   k,
		value: v,
	}
}

func Constructor(cap int) LRUCache {
	l := LRUCache{
		cache:    map[int]*dListNode{},
		head:     initDList(0, 0),
		tail:     initDList(0, 0),
		capacity: cap,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// get查询值 ,没查一次,他就被用一次 ,那么就是被更新 ,需要操作该值,将它移动到链表头部
func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

// 插入一个元素
// 先查有没有 ,没有 init这个节点,然后插入到 头部
//
func (lru *LRUCache) Put(key int, value int) {
	// put值 ,
	if _, ok := lru.cache[key]; !ok {
		// 新建节点
		node := initDList(key, value)
		// 这一步之前漏了 ,至关重要
		lru.cache[key] = node
		// 添加到头
		lru.addToHead(node)
		// 扩大该缓存当前长度
		lru.size++
		// 缓存当前长度当超过缓存容量时
		if lru.size > lru.capacity {
			// 链表中删除该元素
			remove := lru.removeTail()
			// func delete(m map[Type]Type1, key Type)  map源码
			// 该操作用于删除map中该元素
			delete(lru.cache, remove.key)
			// 缓存当前长度缩小
			lru.size--
		}
	} else {
		// 如果查到了
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

// 一些其他操作, 解释双向链表的基本操作,知识现在被放倒 LRU内部
func (lru *LRUCache) addToHead(node *dListNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node

}

func (lru *LRUCache) removeNode(node *dListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *dListNode) {
	// 一定得先删,后添加
	lru.removeNode(node)
	lru.addToHead(node)
}
func (lru *LRUCache) removeTail() *dListNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}

// 手写 hash

func hashMap() {

}
