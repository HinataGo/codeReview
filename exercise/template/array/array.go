package array

// 建立一个结构体保存数组的基本信息,其中包含 数组数据 和 数组长度(数组不是go 的切片没有那么复杂,同时数组的定义必须指明长度)
type Array struct {
	// go1.15没有泛型,使用interface改进
	element []interface{}
	size    int
}

func NewArr(caps int) *Array {
	return &Array{
		element: make([]interface{}, caps),
	}
}

func (a *Array) Caps() int {
	return len(a.element)
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) Empty() bool {
	return a.size == 0
}

// e是需要插入的元素的值
func (a *Array) Add(index int, e interface{}) {
	CheckErr(index, a)
	if a.size == len(a.element) {
		a.resize(2 * a.size)
	}
	// 插入一个元素,需要将后续的数组元素依次向后移动一位
	for i := a.size - 1; i >= index; i-- {
		a.element[i+1] = a.element[i]
	}
	a.element[index] = e
	a.size++
}

func (a *Array) Delete(index int) interface{} {
	// 检查异常
	CheckErr(index, a)
	e := a.element[index]
	// 进行删除操作,所有index 之后的数据向前移动一位
	for i := index + 1; i < a.size; i++ {
		a.element[i-1] = a.element[i]
	}
	// size 缩小, 元素最后一位置为nil
	a.size--
	a.element[a.size] = nil
	// 节约空间的一种优化,简单实现无需写这步
	// if a.size == len(a.element) / 4 && len(a.element) / 2 != 0{
	// 	a.resize(len(a.element) / 2)
	// }

	// 将删除的数据返回出去
	return e
}

// 调整数组大小,事实上go 的数组不像切片会调整大小,go的数组其实是长度不可变的,所以size可变时用slice 而不是array
// 事实上slice 的底层扩容是通过copy做的
func (a *Array) resize(multiple int) {
	a.size = multiple * a.size
}

// 检查异常,索引越界 和 数组溢出
func CheckErr(index int, a *Array) {
	if index < 0 || index > a.size {
		// 此时数组越界,会运行时恐慌,类似于java中error
		panic("add failed, index out of range")
	}
}
