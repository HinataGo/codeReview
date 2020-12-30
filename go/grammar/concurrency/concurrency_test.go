package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// channel 也是第一类对象 和函数一样(可以作为参数,返回值...)

func channel1() {
	c := make(chan int) // 这么定义的双向通道,可以( 收 <- ) 和 (<- 发),单向可以指定,但是没啥用
	go func() {
		for {
			n := <-c // 把c 发送给n
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond) // 不加这行会发现,只输出1
	// 因为第一个1给了channel 之后 输出后就关闭了,1,2都是先收完,然后打印,2还未打印时channel被关闭,所以导致没输出2,
	// 因此,加上休眠,延迟channel关闭,打印输出 2成功
}

func TestChannel(t *testing.T) {
	channel1()
}

// 基于CSP 通过通信来共享内存
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("id %d, revrive &%d \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {

}
