package method

import (
	"fmt"
	"math"
	"testing"
)

// tips :有了函数，为什么还要使用方法？
// Go不是一种纯粹面向对象的编程语言，它不支持类。因此，类型的方法是一种实现类似于类的行为的方法。
// 相同名称的方法可以在不同的类型上定义，而具有相同名称的函数是不允许的。
// 假设我们有一个正方形和圆形的结构。可以在正方形和圆形上定义一个名为Area的方法。这是在下面的程序中完成的。

// go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，
// 接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集
// 方法只是一个函数，它带有一个特殊的接收器类型，它是在func关键字和方法名之间编写的。
// 接收器可以是struct类型或非struct类型。接收方可以在方法内部访问。
// 方法能给用户自定义的类型添加新的行为。它和函数的区别在于方法有一个接收者，给一个函数添加一个接收者，
// 那么它就变成了方法。接收者可以是值接收者，也可以是指针接收者。
// 在调用方法的时候，值类型既可以调用值接收者的方法，也可以调用指针接收者的方法；
// 指针类型既可以调用指针接收者的方法， 也可以调用值接收者的方法。
// 也就是说，不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。

// 定义方法
// func (t Type) methodName(parameter list)(return list) {
//
// }
// func funcName(parameter list)(return list){
//
// }

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
// 指针作为接收者, 可以不加指针,(加指针就是操作原struct ,不加则是复制原struct 在操作,性能差了一些)
// 若不是以指针作为接收者，实际只是获取了一个copy，而不能真正改变接收者的中的数据
func (e *Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d \n", e.name, e.currency, e.salary)
}

// 时方法支持同名定义, 但是函数不行
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func TestCreate(t *testing.T) {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() // Calling displaySalary() method of Employee type

	// 同名方法测试
	// 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
	// method里面可以访问接收者的字段
	// 调用method通过.访问，就像struct里面访问字段一样

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

}

// 测试指针 和 复制 方法区别
type Rectangle1 struct {
	width, height int
}

func (r *Rectangle1) setVal() {
	r.height = 20
}
func (r Rectangle1) setVal1() {
	r.height = 20
}

// 结果差异,pointer 的底层被更改, 复制的则不会
// 20 2
// 2 2
func TestMe(t *testing.T) {
	// pointer
	p := Rectangle1{1, 2}
	s := p
	p.setVal()
	fmt.Println(p.height, s.height)

	// copy
	p1 := Rectangle1{1, 2}
	s1 := p1
	p1.setVal1()
	fmt.Println(p1.height, s1.height)
}

// method是可以继承的，如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method
// 方法是可以继承和重写的
// 存在继承关系时，按照就近原则，进行调用
type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  // 匿名字段, 这里就是继承
	school string
}
type Employee1 struct {
	Human   // 匿名字段
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Employee的method重写Human的method
func (e *Employee1) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

func TestExtend(t *testing.T) {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	mark.SayHi()
	sam := Employee1{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	sam.SayHi()
}
