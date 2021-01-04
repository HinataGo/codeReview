package test21

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 常规思路 两个指针,同时遍历看大小, 依次添加进新的 链表,but 这不是数组不能那么写
	dummy := new(ListNode)
	prev := dummy
	// 设置虚拟头结点 ,建立一个用来推进的指针prev
	// 迭代写法 ,一次性遍历 , 谁小 next指向谁,这样就可以吧值加入到 dummy中
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	// 遍历到最后一个空 ,或者两个都空 ,选择非空的 尾部加入 dummy
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	// 返回dummy的 指针
	return dummy.Next
}
