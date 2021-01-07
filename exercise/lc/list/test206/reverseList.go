package test206

type ListNode struct {
	Val  int
	Next *ListNode
}

// 主要搞清楚 ,数据赋值的顺序
// 链表插入中 我们有 l,r  中间插入 p  那么先将  l.next 给 p  然后 l.next再指向p ,为的是不丢 r

// 这里翻转链表需要一个假的 头 dummy  一开始做尾指针, 然后有一个前遍历到的指针 本身的head
// 这里先建立 第三个 指针 tmp 用于数据交换
// 1. 先把dummy 给  head.next (因为此处已经存了head.next 给 tmp 了 ,所以可以将它重新赋值了)
// 2. 把 head 给dummy 其实就是 dummy 右移
// 3. 将 tmp赋值给head

// 我们发现 1-3 过程中就是把最基本的遍历 list  head = head.next  这一步中间插入 一个 tmp 和dummy
// 从而实现吧 指针 方向调整为反向 最后返回dummy (即新的头结点)

// 策略 : 下次类似题目先写 head = head.next  中间插入两个翻转 ,
// 谁的值先给出去先给谁赋值 ,比如这里,把 head.next 给tmp 那么他就先更新 ,
// 而后发现dummy 又被给出去了 那么再给值给dummy
// 最后到 完成这步骤的遍历

func reverseList(head *ListNode) *ListNode {
	// 翻转链表
	// 把数组的尾指针递归指向前一个指针
	// 新建一个临时节点tmp
	var dummy *ListNode
	for head != nil {
		// 暂存后续链表头结点
		tmp := head.Next
		// 移动的头结点,从头遍历到尾,最后它变成最后的头结点
		head.Next = dummy
		// 移动dummy
		dummy = head
		// 移动cur
		head = tmp
	}
	return dummy
}
