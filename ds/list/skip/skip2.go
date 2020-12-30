package skip

import (
	"math/rand"
	"sync"
)

const MaxLevel = 32 // 8
const SLP = 4

type Node2 struct {
	Forward []Node2
	Value   interface{}
}

func NewNode(v interface{}, level int) *Node2 {
	return &Node2{Value: v, Forward: make([]Node2, level)}
}

type Skiplist struct {
	// 互斥锁，用于保证协程安全
	mutex  *sync.RWMutex
	Header *Node2
	Level  int
}

func Constructor() Skiplist {
	return Skiplist{Level: 1, Header: NewNode(0, MaxLevel)}
}

func (skipList *Skiplist) Add(key int) {
	skipList.mutex.Lock()
	defer skipList.mutex.Unlock()
	update := make(map[int]*Node2)
	node := skipList.Header

	for i := skipList.Level - 1; i >= 0; i-- {
		for {
			if node.Forward[i].Value != nil && node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
			} else {
				break
			}
		}
		update[i] = node
	}

	level := skipList.RandomLevel()
	if level > skipList.Level {
		for i := skipList.Level; i < level; i++ {
			update[i] = skipList.Header
		}
		skipList.Level = level
	}

	newNode := NewNode(key, level)
	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = *newNode
	}
}

func (skipList *Skiplist) RandomLevel() int {
	level := 1
	for (rand.Int31()&0xFFFF)%SLP == 0 {
		level += 1
	}
	if level < MaxLevel {
		return level
	} else {
		return MaxLevel
	}
}

func (skipList *Skiplist) Search(key int) bool {
	skipList.mutex.Lock()
	defer skipList.mutex.Unlock()
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		// fmt.Println("\n Search() Level=", i)
		for {
			if node.Forward[i].Value == nil {
				break
			}

			// fmt.Printf("  %d ", node.Forward[i].Value)
			if node.Forward[i].Value.(int) == key {
				// fmt.Println("\nFound level=", i, " key=", key)
				return true
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}
		} // end for find

	} // end level
	return false
}

func (skipList *Skiplist) Erase(key int) bool {
	skipList.mutex.Lock()
	defer skipList.mutex.Unlock()
	res := false
	update := make(map[int]*Node2)
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		for {

			if node.Forward[i].Value == nil {
				break
			}

			if node.Forward[i].Value.(int) == key {
				// fmt.Println("Remove() level=", i, " key=", key)
				update[i] = node
				res = true
				break
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}

		} // end for find

	} // end level

	for i, v := range update {
		if v == skipList.Header {
			skipList.Level--
		}
		v.Forward[i] = v.Forward[i].Forward[i]
	}
	return res
}

func (skipList *Skiplist) PrintSkipList() {

	// fmt.Println("\nSkipList-------------------------------------------")
	for i := MaxLevel - 1; i >= 0; i-- {

		// fmt.Println("level:", i)
		node := skipList.Header.Forward[i]
		for {
			if node.Value != nil {
				// fmt.Printf("%d ", node.Value.(int))
				node = node.Forward[i]
			} else {
				break
			}
		}
		// fmt.Println("\n--------------------------------------------------------")
	} // end for

	// fmt.Println("Current MaxLevel:", skipList.Level)
}
