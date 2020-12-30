package linked

import "sync"

type node struct{
	// 链表节点的值
	val interface{}
	// 这里实现的跳表是双向的,数组索引代表,第几级索引,可以扩展到n,
	// 需要跳表自动维护
	prev, next []*node
	// skip list 的关键点,用于排序
	key float64
}
type skipList struct {
	dummyHead ,tail *node
	// 代表几级索引
	level  int
	// 记录跳跃表目前包含结点的数量，不包虚拟头结点
	length int
	//
	mutex sync.RWMutex
}
// 定义异常信息
const (
	a = iota
	b
	c
	d
	e
)
// 创建新的节点
func (root *node)NewSkipList() *skipList{
	return &skipList{
		dummyHead: root,
		level: 0,
		length: 0,
	}
}



