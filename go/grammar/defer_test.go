package grammar

import (
	"fmt"
	"testing"
)

// defer 延迟关闭,最上边的 最后关闭
// First, it guarantees that you will never forget to close the file,
// a mistake that's easy to make if you later edit the function
// to add a new return path.
//  Second, it means that the close sits near the open,
//  which is much clearer than placing it at the end of the function.

func TestDefer(t *testing.T) {
	// func trace(s string)   { fmt.Println("entering:", s) }
	// func untrace(s string) { fmt.Println("leaving:", s) }
	//
	// // Use them like this:
	// func a() {
	// 	trace("a")
	// 	defer untrace("a")
	// 	// do something....
	// }
	// 函数本身就有所谓函数栈,
	// 每个调用都是一次压栈,
	// 栈先进后出
	// 随意多个defer也是, 先输出2 后 1
	// 同时不论函数/ 方法内如何实现, defer始终保持 最后执行的操作
	defer fmt.Println(1)
	defer fmt.Println(2)

	// 执行结果
	// entering: b
	// in b
	// entering: a
	// in a
	// leaving: a
	// leaving: b
	// 先压入 defer 只针对了 un ,但是 trace 是一个参数,需要先计算
	// 先计算参数 先打印  entering: b ,返回b (string)
	// 后给了 un("b") ,支持defer参数调入完成
	// 继续执行下一行命令  fmt.Println("in b")
	// 到a() ,先defer 老规矩计算参数值,参数是个函数会先打印数据,然后将返回值交给 un
	// 我们会返现最后un("a") 先输出, un("b")后输出
	// 因此之前的结论正确
	b()
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
