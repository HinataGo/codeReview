package queue

// 最简单的队列
//  像一管道 ，一头进 ，一头出，同时要注意，head不一定非得在出口位置，tail也不一定在入口位置，全看size 和市级的cap轻快那个
type queue struct {
	data []interface{}
	caps int
	head int
	tail int
}

// 初始化时需指明大小
func NewQ(caps int) *queue {
	return new(queue).Init(caps)
}

func (q *queue) Init(size int) *queue {
	q.data = make([]interface{}, size)
	q.caps = size
	q.head = 0
	q.tail = -1
	return q
}

// 向队尾添加，e为待添加值
func (q *queue) Push(e interface{}) bool {
	// 尾指针到达最大尺寸
	if q.tail == q.caps {
		// 队首 为0 ，没有剩余空间，对头在出口，队尾在入口，实在塞不进去了
		if q.head == 0 {
			return false
		} else {
			// 队尾满了，队首到出口还有空隙，搬移数据
			// 就是将数据往出口推， 给队尾到入口腾出空间
			for i := q.head; i < q.tail; i++ {
				q.data[i-q.head] = q.data[i]
			}
			q.tail = q.tail - q.head
			// 对头移动到出口处
			q.head = 0
		}
	}
	q.data[q.tail] = e
	q.tail++
	return true
}

// 对队首弹出元素
func (q *queue) Pop() (interface{}, bool) {
	// 对头队尾在一起，说明没有元素
	if q.head == q.tail {
		return nil, false
	}
	data := q.data[q.head]
	q.data[q.head] = nil
	q.head++
	return data, true
}

// 查看队首值, 不会进行任何删除操作
func (q *queue) Top() interface{} {
	if q.head == q.tail {
		return nil
	}
	data := q.data[q.head]
	return data
}

func (q *queue) qSize() int {
	return q.tail - q.head
}
