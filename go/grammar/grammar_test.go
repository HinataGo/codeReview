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
