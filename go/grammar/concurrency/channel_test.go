package concurrency

// 官网示例,优化channel的反复创建
const MaxOutstanding = 10

// 1. 接收器始终阻塞，直到有数据要接收为止。
// 如果通道未缓冲，则发送方将阻塞，直到接收方收到该值为止。
// 如果通道具有缓冲区，则发送方仅阻塞该值，直到将值复制到缓冲区为止；否则，发送方才阻塞。
// 如果缓冲区已满，则意味着等待直到某些接收器检索到一个值。
//
// 可以像信号灯一样使用缓冲的通道，例如以限制吞吐量。
// 在此示例中，传入的请求被传递到句柄，该句柄将值发送到通道，处理该请求，
// 然后从通道接收一个值，以为下一个使用者准备“信号量”。
// 通道缓冲区的容量限制了同时进行处理的调用数。
var sem = make(chan int, MaxOutstanding)

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func handle(r *Request) {
	sem <- 1   // Wait for active queue to drain.
	process(r) // May take a long time.
	<-sem      // Done; enable next request to run.
}

func Serve1(queue chan *Request) {
	for {
		req := <-queue
		go handle(req) // Don't wait for handle to finish.
	}
}

// 2.这种设计有一个问题：Serve为每个传入的请求创建一个新的goroutine，
// 即使它们中只有MaxOutstanding可以随时运行。
// 如此一来，如果请求太快，程序可能会消耗无限的资源。
// 我们可以通过更改服务以控制goroutine的创建来解决该缺陷。 这是一个显而易见的解决方案，
func Serve2(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func() {
			process(req) // Buggy; see explanation below.
			<-sem
		}()
	}
}

// 3. 错误在于，在Go for循环中，循环变量将在每次迭代中重复使用，
// 因此req变量将在所有goroutine中共享。 那不是我们想要的。
// 我们需要确保每个goroutine的req都是唯一的。 这是一种方法，
// 将req的值作为参数传递给goroutine中的闭包：
func Serve3(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req)
			<-sem
		}(req)
	}
}

// 4. 将此版本与先前版本进行比较，以了解在声明和运行闭包的方式上的差异。
// 另一个解决方案是仅创建一个具有相同名称的新变量，如下例所示：
func Serve4(queue chan *Request) {
	for req := range queue {
		// 为goroutine创建新的req实例
		// Create new instance of req for the goroutine.

		req := req
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

// 这很奇怪
// req := req
// 4.1 但这在Go中是合法且惯用的。 您将获得具有相同名称的变量的新版本，
// 故意在本地隐藏循环变量，但每个goroutine均具有唯一性

// 推荐写法,类似于池化写法
// 5.解决另外一个问题,来一个创建衣蛾goroutine ,是不对的
// 管理资源的方法是启动固定数量的句柄goroutine，这些句柄goroutine全部从请求通道读取。
// goroutine的数量限制了同时处理的数量。
// 该Serve函数还接受一个将告知其退出的通道；
// 启动goroutines后，它将阻止从该通道接收。
func handle5(queue chan *Request) {
	for r := range queue {
		process(r) // 待处理的任务
	}
}

func Serve5(clientRequests chan *Request, quit chan bool) {
	// Start handlers
	for i := 0; i < MaxOutstanding; i++ {
		go handle5(clientRequests)
	}
	<-quit // Wait to be told to exit.
}
