package typesof

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// 这里类型可加可不加,编译器可以自行推断
// 新的golang版本当中，规定const变量也通过驼峰命名法命名，并且首字母必须大写。
const x, y int = 1, 2

func TestConst(t *testing.T) {
	// 这两种定义都可以,但是我们需要注意类型不同,第一个是int32 第二是string类型 ,一个为字符,一个是字符串
	const c = '我'
	const s = "你"
	// reflect.TypeOf() 函数可以查看数据类型 反射机制,性能低,少用
	// 这里还可以使用 %T进行格式化输出类型
	fmt.Println("c type", reflect.TypeOf(c))
	fmt.Println("s type", reflect.TypeOf(s))

	// 常量值可以是编译器能计算出结果的表达式, unsafe len cap
	var a1 [5]int

	const (
		// 这里不写成 := 的简短定义,常量前面已经有const约束了, 而是直接给a进行赋值
		a  = unsafe.Sizeof(c)
		n  = len("hello")
		n2 = cap(a1)
	)
	fmt.Println("a", a)
	fmt.Println("n", n)
	fmt.Println("n2", n2)

	// iota
	const p = iota // p 为0 但是下方iota 会被重置
	const (
		z = 10
		b = 11
		v = 12
		// iota只能在常量下使用 ,这里结果明显 i前面有三个常量.所以i从3开始,如果i在第一个常量位置,则为0
		// iota在const关键字出现时将被重置为0(const内部的第一行之前)，
		// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
		i = iota // 3
		j        // 4
		k        // 5
		m = 100
		q        // 与m相同
		l = iota // 恢复 iota的自增中断 & 必须显示恢复,(C enum会一直自增),因为自增计数会带上前面的,所以为8
	)
	fmt.Println("iota 类似于枚举一个值自增 i,j,k : ", i, j, k)
	fmt.Println("m, q :", m, q)
	fmt.Println("l", l)

	// 常量除了 only read 外和 变量区别?
	// 不优化 分配内存时,常量在预处理时,直接展开,座位指令数据使用
	// other qs
	const x = 5
	const y byte = x // 直接展开x, const y byte = 5
	// const e int = 10 // e定义时 带int (类型),会进行强类型检查
	// const q byte = e // 'p' redeclared in this block
}
