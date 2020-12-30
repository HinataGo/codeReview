package typesof

// 变量定义语法
import (
	"fmt"
	"testing"
)

// 定义，创建了对象并分配内存空间。声明，没有分配空间
// golang不必太纠结直接认为,定义

// 全局变量 ,未使用不报错,var 无需显示初始化, 赋值必须保证左右值相同
var a int
var a1, b1 int

// 不同类型也可以一次性赋值
var a2, b2 = 1, "b2"

// global var not support 简短定义
// a3 := 1

// write test func
func TestVar(t *testing.T) {
	// 简短定义的局部变量必须使用,否则err (1.显示初始化, 2.不能提供数据类型提供也没用 3.智能在函数内部定义)
	// a4 := 4
	a5, a6 := 5, 6
	fmt.Printf("a5 addr %v, a6 value %d \n", &a5, a6) // 输出a5地址 ,a6值
	// 地址不同 ,a被重新定义, 而非重新赋值,简短定义使用需注意
	println("global a: ", &a, a)
	a := 1
	println("part a: ", &a, a)
	// 退化赋值
	// 至少一个新定义的变量, 必须同一定义域
	println("a", &a, a)
	a7, a := 7, 111 // a7为新定义,a为退化赋值(地址不变)
	println("a", &a, a)
	println(a7)
	{
		// 作用于不同这里都是重新定义
		a, a7 := 10, 100
		println(a, a7)
	}
	// 多变量赋值 先计算右值,在赋值,
	x, y := 1, 2
	x, y = y+1, x+2
	println("x", x, "y", y)
}
