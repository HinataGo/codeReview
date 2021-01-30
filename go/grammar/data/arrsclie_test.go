package data

import (
	"fmt"
	"reflect"
	"testing"
)

// array
func TestArr(t *testing.T) {
	a := [...]int{1, 2, 3} // 整型数组,存的元素是整型
	i1 := 1
	i2 := 2
	b := [...]*int{&i1, &i2} // 指针数组 ,存指针的数组
	p0 := &a                 // *[3]int 存整型数组地址的一个指针
	p1 := &b                 // *[2]*int  一个指针,存着指针数组的地址

	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Printf("%T ,%v \n", a, a)
	fmt.Printf("%T ,%v \n", b, b)
	fmt.Printf("%T \n", p0)
	fmt.Printf("%T \n", p1)

	// 与c语言不通, go 的 数组都是值类型,赋值,传参都会进行复制整个数组数据
	// 如果需要避免复制 可以改用指针,或者切片(实际上很少直接用数组,都是用切片)
}

// Arrays are values. Assigning one array to another copies all the elements.
// In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
// The size of an array is part of its type. The types [10]int and [20]int are distinct.
// The value property can be useful but also expensive; if you want C-like behavior and efficiency, you can pass a pointer to the array.
// 在Go和C中，数组的工作方式之间存在主要差异。在Go中， (机翻,大概理解下即可)
//
// 数组是值。 将一个数组分配给另一个数组将复制所有元素。
// 特别是，如果将数组传递给函数，它将接收该数组的副本，而不是指向它的指针。
// 数组的大小是其类型的一部分。 类型[10] int和[20] int是不同的。
// value属性既有用又昂贵。 如果您想要类C的行为和效率，可以将指针传递给数组。

// 说明go 的array 值传递,不存在传值(值拷贝的副本)得指针(引用)
// 数组是带有大小的,eg : [10]int ,数组的大小不同代表不同类型
// 不带数字的[]int 不管怎么弄,都是切片
//  最后一条类C的操作,传指针,而不是传的副本,参考官网示例
/*
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator
*/

// Slices hold references to an underlying array, and if you assign one slice to another,
// both refer to the same array.
// If a function takes a slice argument,
// changes it makes to the elements of the slice will be visible to the caller,
// analogous to passing a pointer to the underlying array.
// A Read function can therefore accept a slice argument rather than a pointer and a count;
// the length within the slice sets an upper limit of how much data to read.

// 切片包含对基础数组的引用，如果将一个切片分配给另一个切片，则两个切片均引用同一数组。
// 如果函数采用slice参数，则对slice的元素所做的更改将对调用者可见，这类似于将指针传递给基础数组
// 因此，Read函数可以接受切片参数，而不是指针和计数。
// 切片内的长度设置了要读取多少数据的上限

// 切片的长度可以更改只要它的基础数组足够长, 只需要从基础的数组分配一部分给切片自己即可
// len 是当前存储的数据长度,cap 是当前切片的最大容量 ,并且针对 nil 的slice 使用合法
// slice 不够时将会自行扩容,然后重新申请一个slice ,将原先的slice 数据拷贝过去,并返回

// 二维切片/ 数组
type Transform [3][3]float64 // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte    // A slice of byte slices.

// 分配方式

/*  1.
// Allocate the top-level slice.
picture := make([][]uint8, YSize) // One row per unit of y.
// Loop over the rows, allocating the slice for each row.
for i := range picture {
	picture[i] = make([]uint8, XSize)
}
*/
/*  2.
// Allocate the top-level slice, the same as before.
picture := make([][]uint8, YSize) // One row per unit of y.
// Allocate one large slice to hold all the pixels.
pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
// Loop over the rows, slicing each row from the front of the remaining pixels slice.
for i := range picture {
	picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
*/
// // Allocate the top-level slice, the same as before.
