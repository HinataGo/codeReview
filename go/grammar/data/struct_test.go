package data

import (
	"fmt"
	"testing"
	"unsafe"
)

// 内存对齐, 直接按照最长的内部元素长度对齐,顺序,会影响 struct 的占用空间-

func TestStruct(t *testing.T) {
	a1 := struct {
		i int  // 对齐宽度8
		a byte // size 1
		c rune // size 4

	}{}
	fmt.Println("size", unsafe.Sizeof(a1)) // i size 8  ,虽然对齐按照最大的,但还是节省空间分配 ,最终a1 size 16
}
