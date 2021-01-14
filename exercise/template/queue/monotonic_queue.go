package queue

// 找出滑动窗口中的最大值/最小值
func mQueue(n int) {
	q := NewQ(n)
	head, tail := 0, -1
	for i := 0; i < n; i++ {
		// check_out 判断条件
		// 对头是否滑出滑动窗口
		for head <= tail && checkOut(q.data[head]) {
			head++
		}
		for head <= tail && check(q.data[tail], i) {
			tail--
		}
		q.data[tail+1] = i
	}
}

func check(i interface{}, i2 int) bool {
	return true
}

func checkOut(i interface{}) bool {
	return true
}
