package typesof

import (
	"fmt"
	"testing"
)

// 类型转换 golang除 常量(const),别名(rune-int32 byte-uint8),未命名( _ )类型外 ,都强制显式类型转换
// golang 不支持运算符重载 ,这样语义更加明确

func TestConversion(t *testing.T) {
	a := 100
	b := byte(a)
	c := a + int(b)
	fmt.Println(c)

	// 不可以将非布尔值作为判断条件
	// if a {
	// }

	// 避免语义歧义, 转换的目标是指针,单向通道,没有返回值的函数类型时,必须使用括号(给所指定的类型加上括号)
	x := 100
	// p := *int(&x) 			// err 指针 解析成 *(int(p))
	// 下面为正确写法
	p := (*int)(&x)
	fmt.Println(p)
	// (func())(x)				//无返回值函数
	// (func()int)(x)			//有返回值的,最好加上方便阅读,可以不加
	// (<-chan int)(c)			//单项通道

}

