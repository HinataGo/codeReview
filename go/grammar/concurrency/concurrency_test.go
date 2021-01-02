package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
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

var ticket = 10
var wg sync.WaitGroup
var mutex1 sync.Mutex // 创建锁头

func sale(s string) {

	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex1.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(s, "售出", ticket)
			ticket--
		} else {
			mutex1.Unlock()
			fmt.Println(s, "all selled")
			break
		}
		mutex1.Unlock()
	}
}
func TestSell(t *testing.T) {
	// 4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。

	wg.Add(4)
	go sale("售票口1")
	go sale("售票口2")
	go sale("售票口3")
	go sale("售票口4")
	wg.Wait()
	// time.Sleep(5.*time.Second)
}
