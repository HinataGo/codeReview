package _go

// 实现信号量 .,并发,控制Connection Pooling 并发数,这里有100个连接,但是我只需要10个
type Semaphore struct {
	c chan struct{}
}

//
func (s Semaphore) run(f func()) {
	s.c <- struct{}{}
	go func() {
		f()
		<-s.c
	}()
}
func Sema(len int) Semaphore {
	return Semaphore{
		c: make(chan struct{}, 10),
	}
}
func main() {
	sema := Sema(10)
	for i := 0; i < 100; i++ {
		sema.run(getXXX)
	}
}
func getXXX() {

}
