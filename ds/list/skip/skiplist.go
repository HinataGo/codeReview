package skip

import (
	"math/rand"
	"sync"
	"time"
)

// 跳表 = 链表 + 有序 + 概率分层
// 参考redis 学习
// 这种skip 是概率型的

const defaultLevel = 32

type SkipListNode struct {
	// key为排序使用的字段,对应元素索引1,2,3...
	key int
	// data为实际存储的数据，可以为任何类型
	data interface{}
	// next为节点指针切片,存储这个节点可到的后继节点 指针 ,存key的地址
	next []*SkipListNode
}

type Skiplist1 struct {
	//  头节点 ,尾节点
	head, tail *SkipListNode
	// 数据总量
	length int
	// 层数,索引是一层一层的
	level int
	// 互斥锁，用于保证协程安全
	mutex *sync.RWMutex
	// 随机数生成器，用于生成随机层数
	random *rand.Rand
}

func (list *Skiplist1) randLev() int {
	lev := 1
	for ; lev < list.level && list.random.Uint32()&0x1 == 1; lev++ {
	}
	return lev
}

func NewSkipList(lev int) *Skiplist1 {
	if lev <= 0 {
		lev = 32
	}
	list := &Skiplist1{}
	list.level = lev
	list.head = &SkipListNode{next: make([]*SkipListNode, lev, lev)}
	list.tail = &SkipListNode{}
	list.mutex = &sync.RWMutex{}
	list.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := range list.head.next {
		list.head.next[index] = list.tail
	}
	return list
}
func constructor() Skiplist1 {
	return *NewSkipList(defaultLevel)
}

// 添加
func (list *Skiplist1) Add(key int, data interface{}) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	// 1.确定插入深度
	level := list.randLev()
	// 2.查找插入部位
	update := make([]*SkipListNode, level, level)
	prev := list.head
	for index := level - 1; index >= 0; index-- {
		for {
			cur := prev.next[index]
			if cur == list.tail || cur.key > key {
				update[index] = prev // 找到一个插入部位
				break
			} else if cur.key == key {
				cur.data = data
				return
			} else {
				prev = cur
			}
		}
	}
	// 3.执行插入
	newNode := &SkipListNode{key, data, make([]*SkipListNode, level, level)}
	for index, node := range update {
		node.next[index], newNode.next[index] = newNode, node.next[index]
	}
	list.length++
}

// 删除
func (list *Skiplist1) Erase(num int) bool {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	// 1.查找删除节点
	node := list.head
	remove := make([]*SkipListNode, list.level, list.level)
	var target *SkipListNode
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > num {
				break
			} else if node1.key == num {
				remove[index] = node // 找到啦
				target = node1
				break
			} else {
				node = node1
			}
		}
	}
	// 2.执行删除
	if target != nil {
		for index, node1 := range remove {
			if node1 != nil {
				node1.next[index] = target.next[index]
			}
		}
		list.length--
		return true
	}
	return false
}

// 查询
func (list *Skiplist1) Search(target int) interface{} {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	node := list.head
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > target {
				break
			} else if node1.key == target {
				return node1.data
			} else {
				node = node1
			}
		}
	}
	return nil
}
func (list *Skiplist1) Length() int {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	return list.length

}
