package typesof

import (
	"fmt"
	"reflect"
	"testing"
)

// 使用make 和new 创建的各种详解
// make 主要针对 slice、map 与 chan 变量类型的内存分配以及相应内部结构的初始化，
// new 可以申请任何类型变量内存，但是拿到内存后会赋零值，他对于 slice、map 以及 chan 类型变量的申请没有太大意义。

func TestMake(t *testing.T) {
	// 创建一个一维切片 ,初值为0,长度为10(长度可以理解为当前存储的元素个数,及当前元素所占的数据大小),元素个数也是10,
	// 容量 20容量才是它真正能存在的数据大小)
	// 区分很简单 看长度有没有定义
	a := make([]int, 10, 20)
	// 目前没有发现类似java初始化同时一次性赋值给所有数据的方法
	// 暂时for遍历赋值
	for i := range a {
		a[i] = 1
		// a[i] += 1
	}

	// var a []int 			// 这种也是切片
	fmt.Println("value a:", a)
	fmt.Println("len a :", len(a))
	fmt.Println("cap a :", cap(a))
	// 输出类型这么输出才对,否则直观上常常无法辨认
	fmt.Println(reflect.TypeOf(a).Kind()) // fmt.Println(reflect.TypeOf(a).Kind().String())

	// b就是一个数组,并且cap和len一样
	var b [10]int
	fmt.Println("value b:", b)
	fmt.Println("len b :", len(b))
	fmt.Println("cap b :", cap(b))
	fmt.Println(reflect.TypeOf(b).Kind())

	c := [10]int{}
	fmt.Println(c)

	// slice
	var s0 []int
	s0 = append(s0, 10)
	fmt.Println(&s0, s0)
	s1 := new([]int)
	*s1 = append(*s1, 10)
	fmt.Println(s1, *s1)
	s2 := make([]int, 0)
	s2 = append(s2, 10)
	fmt.Println(&s2, s2)

	s3 := make([]int, 10)
	s3 = append(s3, 10)
	fmt.Println(&s3, s3)

	// 分配map
	var m0 map[string]int
	fmt.Println(&m0, m0)
	m1 := new(map[string]int)
	fmt.Println(m1, *m1)
	m2 := make(map[string]int, 0)
	m2["key1"] = 100
	fmt.Println(&m2, m2)
	m3 := make(map[string]int, 10)
	m3["key1"] = 100
	fmt.Println(&m3, m3)
	fmt.Println(len(m2))

	// 分配 chan
	var c0 chan int = make(chan int)
	go func(c chan int) {
		c <- 10
	}(c0)

	res := <-c0
	fmt.Println(res)
	s4 := new(string)
	fmt.Printf("%#v", s4)
}
