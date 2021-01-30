package concurrency

import "runtime"

/*
如果可以将计算分解为可以独立执行的独立部分，
则可以并行化计算，并在每个部分完成时发出信号。
*/
// 官网示例
// 1. 不推荐写法
type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // signal that this piece is done
}

const numCPU = 4 // number of CPU cores

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU) // Buffering optional but sensible.
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	// Drain the channel.
	for i := 0; i < numCPU; i++ {
		<-c // wait for one task to complete
	}
	// All done.
}

// 函数runtime.NumCPU返回机器中硬件CPU内核的数量，因此我们可以编写
var numCPU2 = runtime.NumCPU()

// 还有一个功能runtime.GOMAXPROCS，
// 它报告（或设置）用户指定的Go程序可以同时运行的内核数。
// 它的默认值是runtime.NumCPU的值，
// 但是可以通过设置类似命名的shell环境变量或使用正数调用该函数来覆盖
// 用零调用它只是查询值。 因此，如果我们想满足用户的资源请求，我们应该写
var numCPU3 = runtime.GOMAXPROCS(0)

// 确保不要混淆并发的思想（将程序构造为独立执行的组件）和并行性，
// 并发执行并行计算以提高多个CPU的效率。
// 尽管Go的并发特性可以使一些问题易于并行计算，
// 但Go是一种并发语言，而不是并行语言，并且并非所有并行化问题都适合Go的模型。
// https://blog.golang.org/waza-talk
