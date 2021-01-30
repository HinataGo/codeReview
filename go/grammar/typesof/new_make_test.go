package typesof

import (
	"bytes"
	"sync"
	"testing"
)

// go 中有两个原语  用于分配内存资源 new make
// new  分配内存空间,但是不会初始化,是不可以直接用的,只会将其清零
// new(T) return *T ,一个指针
// Since the memory returned by new is zeroed, it's helpful to arrange when designing your data structures
// that the zero value of each type can be used without further initialization.
// This means a user of the data structure can create one with new and get right to work.
func TestNewMake(t *testing.T) {
	// example 使用一种数据结构,创建一个新得数据结构,而不需要进行初始化操作
	type SyncedBuffer struct {
		lock   sync.Mutex
		buffer bytes.Buffer
	}
}

// func NewFile(fd int, name string) *File {
// 	if fd < 0 {
// 		return nil
// 	}
// 	f := File{fd, name, nil, 0}
// 	return &f
// }

// make仅用于 分配并初始化 slice map channel ,并且不会返回指针
// 需要显示调用指针使用new

//  In fact, taking the address of a composite literal allocates a fresh instance each time it is evaluated,
// 所以推荐使用下面的方式
//  return &File{fd, name, nil, 0}

func NewMake() {
	// Unnecessarily complex: 不推荐
	var p *[]int = new([]int) // allocates slice structure; *p == nil; rarely useful
	*p = make([]int, 100, 100)

	// Idiomatic: 推荐

	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints
	v[0] = 1
	v2 := make([]int, 100)
	v2[0] = 1
}
