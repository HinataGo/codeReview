package expression

import (
	"fmt"
	"reflect"
	"testing"
)

// if for  switch 都支持初始化语句

// golang 只有 for 没有while ,没有do while

func TestFor(t *testing.T) {
	// 支持初始化
	for x := 1; x < 10; x++ {
		if x > 5 {
			break
		}
	}

	// 可以代替while,写法多种,看情况选择
	for {
		x := 0
		if x > 10 {
			break
		}
		x++
		break
	}

	a := [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(a).Kind())
	for i, i2 := range a {
		println(&i, &i2) // 多返回,并且定义的参数循环使用 // 这对闭包存在一些问题

	}
	fmt.Println("a data :", a)
	// 说明,表中意思是, i,k = range array  ---> i 是索引,k是array[i] 对应的数
	// range expression 	1st Value 	2nd Value(optional) 	notes
	// array[n]E,*[n]E 		index int 	value E[i]
	// slice[]E 			index int 	value E[i]
	// string abcd 			index int 	rune int 	对于string，range迭代的是Unicode而不是字节，所以返回的值是rune
	// map[k]v 				key k 		value v
	// channel 				element 	none
	// 对于string，range迭代的是Unicode而不是字节，所以返回的值是rune
	// range复制目标数据--- 会直接作用于数组影响数组,可以改用 数组指针,或者切片类型,避免问题
	// 这种从a的复制品中取值, x 值不可变, a[i]是从复制来的数据操作的, 当i== 0时,操作,此时x 已经取了一个值
	for i, x := range a {
		if i == 0 {
			a[0] += 100
			a[1] += 200
			a[2] += 300
		}
		fmt.Println("x : ", x, "a:", a[i])
	}

	// ps : 这里纠正一个很大的误区,很多人认为是深拷贝,浅拷贝,然而golang 中不存在深拷贝,所以不要想错了
	// 仅复制slice,不包括底层数组,此时会操作slice 导致出现后面的x 和 a[i] 值相同
	for i, x := range a[:] {
		// a[0] 输出之前 ,在i == 0 时,x已经取值,所以输出 101,但是后面的可以发现,取值,都是复制slice,不直接作用于数组.
		// 这里我们会发现后面的x还没有取值,因为复制的是slice ,slice可变,在 i == 0时操作完成,后续取x值都变化,出现 x 和 a[i] 值相同
		if i == 0 {
			a[0] += 100
			a[1] += 200
			a[2] += 300
		}
		fmt.Println("x : ", x, ", a:", a[i])
	}

}

func TestIf(t *testing.T) {
	if i := 0; i != 1 {
		i++
	}
}

// 全部case 匹配后才会执行 default语句,所以不必纠结它的位置
// 默认不需要写break, case执行完毕,自行break
// fallthrough , 可以再case执行完后, 强制执行下方的 case ,且该语句必须放在 case块结尾
func TestSwitch(t *testing.T) {
	x := 0
	switch x = 1; x {
	case 1:
		x += 1
		// fallthrough		// 必须卸载case尾部
		if x < 10 {
			break // 这几会break , fallthrough 也不会执行
		}
		fallthrough
	case 2:
		x = 10
	case 3:
		x = 1
		// case x > 0 && x < 10:		// 不能多条件,多条件是 or 关系

	}
	fmt.Println(x) // 因为fallthrough 执行结果 10,case2 被认执行, 但是case3 不会执行
}

func TestLoop(t *testing.T) {
	// goto loop
	// continue
	// break

	// label 定义必须使用
start:
	for i := 0; i < 10; i++ {
		println(i)
		if i > 3 {
			goto start
		}
	}
	// break continue 都可以配合标签使用
}
