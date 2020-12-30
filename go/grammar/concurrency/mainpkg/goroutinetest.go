package main

import (
	"fmt"
	"runtime"
	"time"
)

// easy test
func main() {
	for i := 0; i < 10; i++ {
		// 开了一个协程
		go func(i int) {
			for {
				fmt.Printf("hello %d \n", i)
				runtime.Gosched() // 手动交出资源的控制权,很少使用
			}
		}(i)
	}
	time.Sleep(time.Microsecond) // 微秒
}

// 我们发现它 什么都不打印就退出(不加time.Sleep(time.Microsecond) // 微秒)
// ps :TODO 如果放在test测试问件,结果将会输出,需要探究

// 线程是抢占式调度, 没做完可能强制资源被抢夺, go的协程是非抢占式,主动交出资源,性能更好
// 多个协程可能在一个或多个线程上运行
