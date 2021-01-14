package lru

// Node 链表的节点
type Node struct {
	prev, next *Node

	list *LRUCache2

	key   int
	value int
}

type LRUCache2 struct {
	root *Node // 根节点
	cap  int   // 当前缓存容量
	len  int   // 缓存的长度
}

func Constructor2(capacity int) LRUCache2 {
	l := LRUCache2{
		root: &Node{},
		cap:  capacity,
	}
	l.root.prev = l.root
	l.root.next = l.root
	l.root.list = &l
	return l
}

func (l *LRUCache2) Get(key int) int {
	n := l.get(key)
	if n == nil {
		return -1
	}

	return n.value
}

func (l *LRUCache2) get(key int) *Node {
	for n := l.root.next; n != l.root; n = n.next {
		if n.key == key {
			n.prev.next = n.next
			n.next.prev = n.prev

			n.next = l.root.next
			l.root.next.prev = n
			l.root.next = n
			n.prev = l.root
			return n
		}
	}
	return nil
}

func (l *LRUCache2) Put(key int, value int) {
	n := l.get(key)
	if n != nil {
		n.value = value
		return
	}

	// 缓存满了
	if l.len == l.cap {
		last := l.root.prev
		last.prev.next = l.root
		l.root.prev = last.prev
		last.list = nil
		last.prev = nil
		last.next = nil
		l.len--
	}

	node := &Node{key: key, value: value}
	head := l.root.next
	head.prev = node
	node.next = head
	node.prev = l.root
	l.root.next = node
	l.len++
	node.list = l
}

/**l * Your LRUCache2 object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
