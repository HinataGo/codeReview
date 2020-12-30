package typesof

import (
	"container/list"
	"fmt"
	"testing"
	"unsafe"
)

// 源码详见 src/builtin

/**
u开头代表无符号
类型				长度		默认值		备注
bool			1		false
byte			1		0			uint8
int,uint		4,8		0			默认类型,根据系统32/64(本机为64,则长度都为8)
int8,uint8  	1		0			-128--127, 0--255
int16,uint16	2		0			-32768--32767,0--65535
int32,uint32	4		0
int64,uint64	8		0
float32			4		0.0			默认类型
float64			8		0.0
complex64		8
complex128		16
rune			4		0			Unicode code point(Unicode字符)对应 int32
uintptr			4,8		0			存储指针的uint(64位机器上 8)
string					""			(这里是空字符串),不是NULL
array
struct
function				nil			函数
interface				nil			接口
map						nil			字典,引用类型
slice					nil			切片,引用类型
channel					nil			通道,引用类型
list					nil			链表 container包里,双向链表


ps: 别名(类型一致无需转换)
rune 	alias for int32
byte 	alias for int8
这里提一下64位下 int 和int64 底层结构相同并不属于别名,他们不能直接相互赋值使用
*/

func TestBasic(t *testing.T) {
	var a int
	var b uint
	fmt.Println("int size 64", unsafe.Sizeof(a))
	fmt.Println("uint size 64", unsafe.Sizeof(b))
	var p uintptr
	fmt.Println("uintptr ize 64", unsafe.Sizeof(p))
	var s rune = ' '
	fmt.Printf("s value %c \n", s)
	fmt.Printf("s type %T \n", s)
	l := list.New()
	fmt.Printf("list default type %T \n", l)
}

// 引用类型测试必须使用make
func TestReference(t *testing.T) {
	// slice
	m := make([]int, 0, 10)
	m = append(m, 1)
	fmt.Println(m)
	// map
	// 分配内存,但不完整创建,(创建字典类型包装/ 指针包装, 实际上所需的内存,并没有分配剑指存储内存,也没初始化散列桶等内部属性)
	// m2 := new(map[string]int)
	// p := *m2
	// p["a"] = 1
	// fmt.Println(p)
	m3 := make(map[string]int)
	m3["a"] = 1
	fmt.Println(m3)
	// channel

}
