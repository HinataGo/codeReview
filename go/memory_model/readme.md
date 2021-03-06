# Go 内存模型 (Memory model)
## 官翻
###  介绍
- Go内存模型指定了一种条件,在该条件下可以保证 在一个 goroutine中读取一个变量a,可以观察到不同 goroutine中写入同一变量所产生的值
### tips
- 修改由 多个goroutine 同时访问 的数据的程序时 必须序列化此类访问。
- 要进行序列化访问时, 使用channel / 其他的额同步原语 比如 sync | sync/atomic ,来保护数据

### 执行之前, 关于重排序的一个问题(比如 CPU就有乱序执行,后还有一个重排序过程)
- tips
  - 会在一条指令执行过程中（比如去内存读数据，大概慢100多倍），去同时执行另一条指令，前提是两条指令没有依赖关系
- 单个 go程中, 读 read / 写 write 是必须按照程序指定的顺序 执行一样
- 因此对于上面的解释就是: 仅当重新排序不会改变语言规范所定义的该goroutine中的行为时，编译器和处理器才可以对单个goroutine中执行的读取和写入进行重新排序
- 由于此重新排序，一个goroutine观察到的执行顺序可能与另一个goroutine察觉到的执行顺序不同

- 这里还有一个问题及时, Go程序关于执行顺序的定义
    - 1. 如果事件e1发生在事件e2之前，那么我们说e2发生在e1之后
    - 2. 如果 e1 在 e2 之前发生 并且 e2 之后 e1 前后没有任何改变,那么认为 e1 ,e2 是同时执行的
- Go中定义,满足下面两个条件, 允许 对 变量v 的 写过程进行观察(r 读)
    - 1.1 r 在 w 之前不会发生
    - 1.2 在w 之后 但在 r 之前没有发生对 v 的其他 写操作
- 还有一种: 为了保证 对于变量v 的读取 时可以观察到 退役变量v 的特定的写入操作(w),并且明确 w 是唯一允许r 官产的写操作
    - 1. w 发生在 r 之前
    - 2. 对于共享变量 v 的任何其他写擦走 都发生在 w之前 或者 r同时发生
- 单个goroutine 没有并发性, 因此这两个定义等价
- 大于单个机器字的值的读取和写入将以未指定的顺序充当多个机器字大小的操作
###  同步化  Synchronization
#### 程序初始化 init
- 在单个goroutine中运行，但是该goroutine可能会创建其他同时运行的goroutine。
- 如果包p导入了包q，则q的init函数的完成发生在任何p的开始之前。
- 函数main.main的启动发生在所有init函数完成之后。
#### Goroutine 创建
- 在需要执行的 func 之前加  go 即可
#### Goroutine 破坏 (数据竞争)
```bazaar
var a = "test"

func main() {
	var a string

	go func() { a = "hello" }()
	fmt.Println(a)

}
// 这段输出 test 而 匿名函数不会被执行
// fmt.Print 输出的a 输出的时候没有看到 a被go修改过 这是出现了数据竞争
```
- 0. 操作前加锁,用channel  杜绝一切可能导致 data race(而且你这个可能出现并发读写的问题, )
- 1. 分配给a不会跟随任何同步事件，因此不能保证任何其他goroutine都会遵守该事件。 实际上，积极的编译器可能会删除整个go语句
- 2. 如果必须通过另一个goroutine来观察goroutine的影响， 请使用同步机制（例如锁定或通道通信）来建立相对顺序
    
#### channel 通信 (结局上述问题)
- 通常在不同的goroutine中，将特定通道上的每个发送与该通道上的相应接收进行匹配。
- 通道上的发送发生在该通道上的相应接收完成之前。
```bazaar
// 该程序（如上所述，但是交换了send和receive语句并使用了未缓冲的通道）
// 也保证打印“ hello，world”。
// 对a的写操作发生在c上的接收之前，
// 发生在相应的c上发送完成之前，发生在打印之前。
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```
- 如果通道已缓冲,则不能保证程序会打印“ hello，world”。
  - 例如，c = make（chan int，1） 它可能会打印空字符串，崩溃或执行其他操作。
- 

## 个人实践学习