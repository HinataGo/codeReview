package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// 和actor直接同不同, CSP 通过Channel 进行通讯,更松耦合
// Go中的Channel 是有容量限制并且独立于处理Goroutine ,Erlang Actor中mailbox无限,接受信息也是被动的

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	// time.Sleep(time.Second * 1)
}
