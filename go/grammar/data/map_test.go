package data

import (
	"fmt"
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := make(map[string]int, 10) // size / len 10 , key string ,value int

	m1["a"] = 1
	m1["a"]++ // m1["a"]  = m1["a"] + 1
	m2 := make(map[string]string)
	m2["a"] = "1"
	// m2["a"]++ 		// errors Invalid operation: m2["a"]++ (non-numeric type string)
	fmt.Println(m1)
	// 迭代期间 删除 新增键值是安全的
	// TODO 数据竞争,解决一个任务占用写操作,其他任务无法对map执行并发操作(读,写,删除)
}

// map 支持很多类型
// 键可以是定义了相等运算符的任何类型，
// 例如整数，浮点数和复数，字符串，指针，接口（只要动态类型支持相等），结构和数组。
// 切片不能用作映射键，因为未在其上定义相等性。
// 像切片一样，映射保留对基础数据结构的引用。
// 可以使用带冒号 双引号 构建 或者初始化
var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
	"a":   10,
}

// 分配和获取映射值的方式类似于对数组和切片执行相同的操作，不同之处在于索引不必为整数
// offset := timeZone["EST"]

// map[key] 查到即为数据, 查不到返回0,可以添加bool判断,语义清晰  ok 查到返回true,查不到返回 false
var seconds int
var ok bool

// seconds, ok = timeZone[tz]

// 常用操作
func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		fmt.Println(ok)
		return seconds
	}
	fmt.Println(timeZone[tz])
	fmt.Println(ok)
	log.Println("unknown time zone:", tz)
	return 0
}

func TestMap1(t *testing.T) {
	offset("nil")
}
