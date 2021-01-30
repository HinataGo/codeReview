package _interface

import (
	"fmt"
	"net/http"
	"os"
	"sort"
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

// 自定义打印机可以通过String方法实现，而Fprintf可以使用Write方法生成任何内容的输出。
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
// func (s Sequence) String() string {
// 	s = s.Copy() // Make a copy; don't overwrite argument.
// 	sort.Sort(s)
// 	str := "["
// 	for i, elem := range s { // Loop is O(N²); will fix that in next example.
// 		if i > 0 {
// 			str += " "
// 		}
// 		str += fmt.Sprint(elem)
// 	}
// 	return str + "]"
// }

// 上面那个复杂度太高,这个是修改版
// Go程序中，习惯用法是转换表达式的类型以访问不同的方法集
func (s Sequence) String() string {
	s = s.Copy()

	// sort.Sort(s)
	// return fmt.Sprint([]int(s))

	// IntSlice将Interface的方法附加到[] int，并按升序排序
	// 和上面功能类似
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

// 在实践中，这种情况较不常见，但可以有效。
func TestS(t *testing.T) {
	var arr []int
	arr = []int{1, 4, 2}
	v := Sequence.String(arr)
	fmt.Println(v)
}

// Simple counter server.
type Counter struct {
	n int
}

// 在真实服务器中，访问ctr.n需要防止并发访问。 请参阅sync和atomic软件包以获取建议。
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

// 但是为什么要把Counter作为一个结构呢？ 整数就足够了。
// （接收方必须是一个指针，这样增量才能对调用方可见。）
// Simpler counter server.
type Counter1 int

func (ctr *Counter1) ServeHTTP2(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

// 如果您的程序有一些内部状态需要通知已访问页面怎么办？
// 将频道绑定到网页。
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
// 每次访问都会发送通知的渠道。
// 可能希望通道被缓冲
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}
func TestHttp1(t *testing.T) {
	http.Handle("/args", http.HandlerFunc(ArgServer))
}
