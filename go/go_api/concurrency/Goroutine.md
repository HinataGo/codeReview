# Goroutine

- 普通函数
    - 线程是 main调度 --> do work(单项通信,一层层调用)
    - 协程 main <--> do work (双向通信,权限资源可以互相给,多个协程可以再一个线程,也可以多个线程之间) (主流,新版java也支持协程了), 
    
- go语言 goroutine 协程
    - 直接加 go 就可以发送给调度器运行,被认作是 协程
    - 不需要定义时区分是否是异步函数
    - 调度器在合适点有调度器进行切换(go 不需要自己手动切换)
    - race检测数据访问冲突
- ps: 子程序(Subroutine)是协程(Coroutines)的一个特例
- goroutine 切换点(仅参考,还有其他情况)
    - IO ,select
    - channel
    - 等待锁
    - 有时候,函数调用
    - runtime.Gosched()
    
## channel 看concurrency_test.go代码
- 接受方定义 a  <-chan
- 发送方定义 b  chan<-

