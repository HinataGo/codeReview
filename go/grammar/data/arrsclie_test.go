package data

import (
	"fmt"
	"reflect"
	"testing"
)

// array
func TestArr(t *testing.T) {
	a := [...]int{1, 2, 3} // 整型数组,存的元素是整型
	i1 := 1
	i2 := 2
	b := [...]*int{&i1, &i2} // 指针数组 ,存指针的数组
	p0 := &a                 // *[3]int 存整型数组地址的一个指针
	p1 := &b                 // *[2]*int  一个指针,存着指针数组的地址

	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Printf("%T ,%v \n", a, a)
	fmt.Printf("%T ,%v \n", b, b)
	fmt.Printf("%T \n", p0)
	fmt.Printf("%T \n", p1)

	// 与c语言不通, go 的 数组都是值类型,赋值,传参都会进行复制整个数组数据
	// 如果需要避免复制 可以改用指针,或者切片(实际上很少直接用数组,都是用切片)
}
