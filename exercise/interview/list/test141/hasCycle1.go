package test141

// 环形链表 ，这里是时空间复杂度为 O(n) ,题目要求O(1)空间
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	res := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := res[head]; ok {
			return true
		} else {
			res[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}
