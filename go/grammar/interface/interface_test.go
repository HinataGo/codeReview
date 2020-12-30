package _interface

import (
	"fmt"
	"testing"
)

// 定义接口的方法
// /* 定义接口 */
// type interface_name interface {
//   method_name1 [return_type]
//   method_name2 [return_type]
//   method_name3 [return_type]
//   ...
//   method_namen [return_type]
// }
//
// /* 定义结构体 */
// type struct_name struct {
//   /* variables */
// }
//
// /* 实现接口方法 */
// func (struct_name_variable struct_name) method_name1() [return_type] {
//   /* 方法实现 */
// }
// ...
// func (struct_name_variable struct_name) method_namen() [return_type] {
//   /* 方法实现*/
// }

// interface可以被任意的对象实现
// 一个对象可以实现任意多个interface
// 任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface

type Phone interface {
	call() // 这是一个方法,定义在接口里面
}

type NokiaPhone struct {
}

// 这是一个接口
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func TestInterface(t *testing.T) {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

// interface & Duck Typing
// go 语言作为一门静态语言，它通过通过接口的方式完美支持鸭子类型。
// 在静态语言如 Java, C++ 中，必须要显示地声明实现了某个接口，之后，才能用在任何需要这个接口的地方。
// 如果你在程序中调用某个数，却传入了一个根本就没有实现另一个的类型，那在编译阶段就不会通过。
// 这也是静态语言比动态语言更安全的原因。
// go 采用了折中的做法：不要求类型显示地声明实现了某个接口，只要实现了相关的方法即可，编译器就能检测到。

// Go语言的多态性：接口可以在Go中隐式地实现。如果类型为接口中声明的所有方法提供了定义，则实现一个接口。让我们看看在接口的帮助下如何实现多态。
// 任何定义接口所有方法的类型都被称为隐式地实现该接口。

/*
// 安全类型断言
<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 )
//非安全类型断言
<目标类型的值> := <表达式>.( 目标类型 )

*/
type student struct {
}

func TestAssertion(t *testing.T) {
	// var i1 interface{} = new(Student)
	// s := i1.(Student) // 不安全，如果断言失败，会直接panic
	//
	// fmt.Println(s)

	var i2 interface{} = new(student)
	s, ok := i2.(student) // 安全，断言失败，也不会panic，只是ok的值为false
	if ok {
		fmt.Println(s)
	}
	// 接口对象不能调用接口实现对象的属性
	// 断言其实还有另一种形式，就是用在利用 switch语句判断接口的类型。 case 语句的顺序是很重要的，因为很有可能会有多个 case匹配的情况
	//  s.(type) 必须在switch case 中  // use of .(type) outside type switch

	// switch ins := s1.(type)
	// Invalid type switch guard: ins := s1.(type) (non-interface type student on left)
	// 传进来的参数不是interface类型那么做类型断言都是回报 non-interface的错误
	// 需要借助interface{}()进行转换
	s1 := interface{}(s)

	switch ins := s1.(type) {
	case string:
		fmt.Println("三角形。。。", ins)
	case rune:
		fmt.Println("圆形。。。。")
	default:
		fmt.Println("整型数据。。")
	}

}
