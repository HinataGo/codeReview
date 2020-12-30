package data

import (
	"fmt"
	"testing"
	"unsafe"
)

// string 是一个复合结构
// type stringStruct struct {
//	str unsafe.Pointer

//	len int
// }
// 默认存储UTF8编码的Unicode字符
// 默认值 "" 不是  nil
// TODO for 遍历 字符串有 byte 和 rune

//

// TODO string默认不能直接修改,非要修改,先转换类型为 []rune/ []byte,完成后再转换成string ,但是不论如何转换,都须重新分配内存,并复制数据
// 转换操作有时候性能较差,可以使用unsafe 方法改善
func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
func TestString(t *testing.T) {
	ss := []byte("hello")
	s := toString(ss)
	// 可以看书地址是一致的, 避免了底层数组的复制,很多web framework 中都使用这种,但是相对的 unsafe方法有不安全的隐患
	// 同样编译器在部分场合会进行优化性能,避免额外分配和复制操作 : 将[]byte转换为 string key,在map[string] 查询的时候
	// 将string 转换为byte[] 进行for range 迭代的时候,直接取字节赋值给局部变量
	// 打印地址建议使用 %p
	// 详见 https://studygolang.com/articles/2644
	fmt.Printf("ss: %p \n", &ss)
	fmt.Printf("s : %p \n", &s)
}
