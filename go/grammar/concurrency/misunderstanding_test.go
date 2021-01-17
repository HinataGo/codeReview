package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// fatal error: all goroutines are asleep - deadlock

// 对于同一无缓冲通道，在接收者未准备好之前，发送操作是阻塞的。而此处的通道ch就是缺少一个配对的接收者，因此造成了死锁。

func TestCh1(t *testing.T) {
	ch := make(chan string, 1)
	go func() {
		ch <- "hello world"
	}()
	fmt.Println(<-ch)
}

// 解决方法
// 并发的方式调用匿名函数func   go func(){} //go中的表达式必须是函数调用
// 第一种添加配对的接收者；
// go func() {
// 	ch <- "hello world"
// }()

// 第二种将默认的通道替换成缓冲通道
// ch := make(chan string, 1)

// go 执行类似defer ，会先进行求值，并且没有接收端 的原因，会发生卡死，system会判断 deadlock（go程死锁）

func TestGo(t *testing.T) {
	ch := make(chan int, 10)
	fmt.Println(<-ch)
	ch <- 10
	time.Sleep(1 * time.Second)

}

// 同样的添加配对的接受者即可,这里后面需要加()
// go func() {
//		fmt.Println(<-ch)
//	}()
