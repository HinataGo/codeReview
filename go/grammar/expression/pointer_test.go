package expression

import "testing"

// 内存地址:内存单元的唯一编号
// 指针是一个实体,  指针会 分配内存空间 (就是一个专门存储内存地址的 整型变量)
// 内存地址和额指针不同
func TestPointer(t *testing.T) {
	// & 取地址运算符,
	// *指针运算符, 简介引用目标对象,定义指针
	x := 10
	var p *int = &x
	*p += 20
	println(p, *p)
}
