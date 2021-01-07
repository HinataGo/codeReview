package skip

// 说明,这不是跳表实现,这是leetcode取巧的做法 ,因为hash 的存取也是极其快的,所以这里直接用map做list的
type Skiplist4 struct {
	M map[int]int
}

func Constructor4() Skiplist4 {
	s := Skiplist4{}
	s.M = map[int]int{}
	return s
}

func (s *Skiplist4) Search(target int) bool {
	return s.M[target] > 0
}

func (s *Skiplist4) Add(num int) {
	s.M[num]++
}

func (s *Skiplist4) Erase(num int) bool {
	if s.M[num] > 0 {
		s.M[num]--
		return true
	}
	return false
}
