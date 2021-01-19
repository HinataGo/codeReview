package data

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := make(map[string]int, 10) // size / len 10 , key string ,value int

	m1["a"] = 1
	m1["a"]++ // m1["a"]  = m1["a"] + 1
	m2 := make(map[string]string)
	m2["a"] = "1"
	// m2["a"]++ 		// error Invalid operation: m2["a"]++ (non-numeric type string)
	fmt.Println(m1)
	// 迭代期间 删除 新增键值是安全的
	// TODO 数据竞争,解决一个任务占用写操作,其他任务无法对map执行并发操作(读,写,删除)
}
