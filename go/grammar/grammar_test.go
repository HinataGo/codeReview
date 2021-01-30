package grammar

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// easy test
func TestDemo01(t *testing.T) {
	t.Log("test log")
}

// if
func TestOther(t *testing.T) {
	if _, err := aa(); err != nil {
		t.Log("OK")
	} else {
		t.Log("no")
	}
}
func aa() (int, *int) {
	return 0, nil
}

// switch

// There is no automatic fall through, but cases can be presented in comma-separated lists.
func shouldEscape(c byte) bool {
	switch c {

	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch i {
		case 0, 1:
			t.Log("A")
		case 2, 3:
			t.Log("B")
		case 4, 5, 6:
			t.Log("C")
		}

	}
	fmt.Println(shouldEscape('a'))

	var tp interface{}
	switch tp := tp.(type) {
	default:
		fmt.Printf("unexpected type %T\n", tp) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", tp) // t has type bool
	case int:
		fmt.Printf("integer %d\n", tp) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *tp) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *tp) // t has type *int
	}
}

// loop label switch continue
// break可以指定跳跃出 指定label ,另外一种break 只会跳出switch 不会跳出for
// continue 只能在 loop中使用
func TestLoop(他 *testing.T) {
	var src []int
	src = []int{1, 2, 3}
Loop:
	for n := 0; n < len(src); n += 1 {
		switch {

		case src[n] < 3:
			//
			if src[n] > 0 {
				fmt.Println(n)
				break Loop
			}
			if src[n] > 1 {
				fmt.Println(n)
				break
			}

		}
	}
	src = []int{1, 3, 5, 10, 6, 7, 9}
	a := -1

	// continue 后面加不加label作用一致,所以 continue作用很显而易见了,5 不输出
loop2:
	for i := 0; i < len(src); i++ {

		if src[i] == 5 {
			continue loop2
		}
		a = src[i]
		fmt.Println(a)
	}
}

// array slice
func TestArr(t *testing.T) {
	// 声明, 这里a, b类型认为是不一样的因为 a数组未指定长度默认0
	var a []int
	var b [10]int
	fmt.Print(a, b)
}

// map 字典
func TestMap(t *testing.T) {
	m := map[int]int{1: 4, 2: 5, 3: 6}
	t.Log(m[1])
	m2 := make(map[int]int, 2)
	t.Logf("len m2 = %d", len(m2))

}

// 使用map实现myset
func TestMyset(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] {
		t.Log("yes")
	}
}

// 函数
type ss struct {
	name string
	id   int
}

// 多返回值
func testFun(t *testing.T) (int, int) {
	return 0, 1
}

// 函数可以接受结构体接受者,这种称作方法
func (s *ss) fun1(a int) {

}

// string  ,另外一个常用字符串函数包 strconv包
func TestString(t *testing.T) {
	var s1 string = "qqqq"
	fmt.Println(s1[1])
	// s1[1]  = 's' Cannot assign to s1[1]
	// string 是不可变的byte slice

	// 字符串分割
	s2 := "a,b,c"
	p := strings.Split(s2, ",")
	for _, e := range p {
		t.Log(e)
	}
	t.Log(strings.Join(p, "-"))

	//strconv
	s3 := strconv.Itoa(10)
	t.Log("str" + s3)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(2 + i)
	}
}

// ps Unicode是一种字符集 ,UTF8(最常用)是它的一种存储实现(转换为字节序列的规则),诸如UTF16都是Unicode编码

// go都是传值,没有传引用,没有所谓深拷贝

// 可变参数
func TestVar(t *testing.T) {

}

// go闭包 Closure
// 以goroutines运行的闭包会发生 错误
func Closure() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

// 可能会错误地期望看到a，b，c作为输出。 您可能会看到的是c，c，c。
// 这是因为循环的每次迭代都使用变量v的相同实例，因此每个闭包都共享该单个变量。
// 当闭包运行时，它将在执行fmt.Println时打印v的值，但是自启动goroutine以来，
// v可能已被修改。 为了帮助在此问题和其他问题发生之前发现它们，请运行go vet。
//
// 要将v的当前值绑定到每个闭包启动时，必须修改内部循环以在每次迭代时创建一个新变量。
// 一种方法是将变量作为参数传递给闭包：
// 在此示例中，v的值作为参数传递给匿名函数。
// 然后可以在函数内部将该值作为变量u访问。
func Closure2() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

// 使用声明样式似乎更奇怪，但是在Go中可以正常工作，这甚至更容易：
// 这种语言的行为（未为每次迭代定义新的变量）可能是一个错误。
// 可能会在更高版本中解决，但出于兼容性考虑，在Go版本1中无法更改。
func Closure3() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v // create a new 'v'.
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
