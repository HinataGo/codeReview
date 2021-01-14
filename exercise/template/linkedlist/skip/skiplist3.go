package skip

import (
	"math"
	"math/rand"
)

// right + down 版本skipList

const (
	maxLevel = 16
	maxRand  = 65535
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

// SkipList3
type SkipList3 struct {
	head *Node3
}

// Node3
type Node3 struct {
	right *Node3
	down  *Node3
	val   int
}

func Constructor3() SkipList3 {
	left := make([]*Node3, maxLevel)
	right := make([]*Node3, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &Node3{val: math.MinInt16}
		right[i] = &Node3{val: math.MaxInt16}
	}
	// 第0层为最下层
	for i := 0; i < maxLevel; i++ {
		left[i].right = right[i]
		if i != 0 {
			left[i].down = left[i-1]
			right[i].down = right[i-1]
		}

	}
	return SkipList3{left[maxLevel-1]}
}

func (this *SkipList3) Search(target int) bool {
	head := this.head
	for head != nil {
		if head.right != nil {
			if head.right.val < target {
				head = head.right
			} else if head.right.val == target {
				return true
			} else {
				head = head.down
			}
		} else {
			head = head.down
		}
	}
	return false
}

func (this *SkipList3) Add(num int) {
	pre := make([]*Node3, maxLevel)
	head := this.head
	// 找到每次最小于num的node
	for lv := maxLevel; head != nil; head = head.down {
		for head.right != nil && head.right.val < num {
			head = head.right
		}
		lv--
		pre[lv] = head
	}

	n := randLevel()
	arr := make([]*Node3, n)
	tmp := &Node3{val: math.MinInt16}
	// 从上面开始创建
	for i := n - 1; i >= 0; i-- {
		a := arr[i]
		p := pre[i]
		a = &Node3{val: num, right: p.right}
		p.right = a
		// 0层不应该有向下指针
		if i != n-1 {
			tmp.down = a
		}
		tmp = a
	}
}

func (this *SkipList3) Erase(num int) bool {
	ret := false
	head := this.head
	for head != nil {
		if head.right != nil {
			if head.right.val < num {
				head = head.right
			} else if head.right.val == num {
				head.right = head.right.right
				head = head.down
				ret = true
			} else {
				head = head.down
			}
		} else {
			head = head.down
		}
	}
	return ret
}
