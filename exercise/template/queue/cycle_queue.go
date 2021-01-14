package queue

type CycleQueue struct {
	data  []interface{} // 存储空间
	front int           // 前指针,前指针负责弹出数据移动
	rear  int           // 尾指针,后指针负责添加数据移动
	cap   int           // 设置切片最大容量
}

func NewCQ(cap int) *CycleQueue {
	return &CycleQueue{
		data:  make([]interface{}, cap),
		cap:   cap,
		front: 0,
		rear:  0,
	}
}

// 入队操作
// 判断队列是否队满,队满则不允许添加数据
func (q *CycleQueue) Push(data interface{}) bool {
	// check queue is full
	if (q.rear+1)%q.cap == q.front { // 队列已满时，不执行入队操作
		return false
	}
	q.data[q.rear] = data         // 将元素放入队列尾部
	q.rear = (q.rear + 1) % q.cap // 尾部元素指向下一个空间位置,取模运算保证了索引不越界（余数一定小于除数）
	return true
}

// 出队操作
// 需要考虑: 队列为空没有数据返回了
func (q *CycleQueue) Pop() interface{} {
	if q.rear == q.front {
		return nil
	}
	data := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.cap
	return data
}

// 因为是循环队列, 后指针减去前指针 加上最大值, 然后与最大值 取余
func (q *CycleQueue) QueueLength() int {
	return (q.rear - q.front + q.cap) % q.cap
}
