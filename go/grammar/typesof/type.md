# 辨析

### 1. type
- 注意重新定义类型是 没有 = ，此时认为两个不是同一个类型
- 但是存在 = ，起别名，本质上还是一致的
```bazaar
// 基于int创建新类型
type MyInt1 int
// 起别名
type MyInt2 = int

func main() {
	var i int =0
	// 错误， type 后的类型  MyInt1 默认与 int不是同一类型
	// 将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1,i2)
}

```

### 2.
```bazaar
// error 返回类型要么都显示的表达返回的变量名 ，要么都不写
// (sum int, error) 去掉sum  或者给error 补充一个
func funcMui(x,y int)(sum int, error){
	return x+y,nil
}
```

### 3. append & new & make
- 推荐 make初始化不使用new,同时make需要指明大小
- make(type, len, cap) 其中cap可以省略，省略后cap默认和len一致
- ps如果是切片可以自己扩容，数组不可以
```bazaar
# 不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
# Cannot use 's' (type *[]int) as type []Type
func main() {
	s := new([]int)
	s = append(s, 1)
	fmt.Println(s)
}
```
### 4. var & 简写定义
```bazaar
// var 内部不要写 简写定义, 即 :=
var(
    size := 1024
    max_size = size*2
)

func main() {
    fmt.Println(size,max_size)
}
```
### 5. const
```bazaar
// const 内部支持 _ 定义
const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

```
### 6. slice
```bazaar
// s2 不能做 int append 无法直接将一个切屏append到另外一个slice 切片 
// 注意这里[]int 是 切片 不是数组
    s1 := []int{1, 2, 3}
	s2 := []int{4, 5 ,0}
	s1 = append(s1, s2)
	
// 操作符 [i,j]。基于数组（切片）可以使用操作符 [i,j] 创建新的切片，从索引 i，到索引 j 结束，截取已有数组（切片）的任意部分，返回新的切片，新切片的值包含原数组（切片）的 i 索引的值，但是不包含 j 索引的值。i、j 都是可选的，i 如果省略，默认是 0，j 如果省略，默认是原数组（切片）的长度。i、j 都不能超过这个长度值。

// 假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：长度：j-i，容量：k-i。

*// 截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。所以例子中，切片 t 为 [4]，长度和容量都是 1。
```

### 7. struct
- 可比较的，常见的有 bool、数值型、字符、指针、数组等，
- 像切片、map、函数等是不能比较的。
```bazaar

func main() {
    sn1 := struct {
        age  int
        name string
    }{age: 11, name: "qq"}
#    sn3 和 sn1 不可以比较 ，内部顺序不同
    sn3:= struct {
            name string
            age  int
        }{age:11,name:"qq"}
    sn2 := struct {
        age  int
        name string
    }{age: 11, name: "qq"}

    if sn1 == sn2 {
        fmt.Println("sn1 == sn2")
    }

    sm1 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}
    sm2 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}

    if sm1 == sm2 {
        fmt.Println("sm1 == sm2")
    }
# 结构体只能比较是否相等，但是不能比较大小。
# 相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体；
# 如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
}
```

### 8. nil
```bazaar
#Cannot assign nil without explicit type
# 正确做法在 x 后面加 var x interface{}= nil / var x error = nil
var x = nil
// nil 值。nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。强调下 var x error = nil 选项的 error 类型，它是一种内置接口类型，看下方贴出的源码就知道，所以 var x error = nil 是对的。
// 源码
type error interface {
    Error() string
}
```

### 9. slice [low : high : max]
- low起点 
- high终点(不包括此下标)
- cap = max-low, 容量
- len = high-low ,长度
```bazaar
    a := [5]int{1, 2, 3, 4, 5}
	t := a[2:4:5]
	fmt.Println(t)         // [3 ,4]
	fmt.Println(len(t))     // 2
	fmt.Println(cap(t))     // 3
	
	
```

### 10. channel
```bazaar
    var ch chan int
	// 写channel 必须带上值
	ch <- 1
	// 读channel，可以不带值
	<- ch
```
### 11、range
```bazaar
func main() {

     slice := []int{0,1,2,3}
     m := make(map[int]*int)

     for key,val := range slice {
         m[key] = &val
     }

    for k,v := range m {
        fmt.Println(k,"->",*v)
    }
}
// for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.
// 修改 
for key,val := range slice {
         value := val
         m[key] = &value
     }
```
### 12.map
```bazaar
type person struct {  
    name string
}

func main() {  
    var m map[person]int
    p := person{"mike"}
    fmt.Println(m[p])
}
// 打印一个 map 中不存在的值时，返回元素类型的零值
// m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0
```

### 13.
```bazaar
// 两个不同类型的数值不能相加，编译报错

```

### 14.函数赋值
```bazaar
// h := hello 将 hello() 赋值给变量 h ,h 值不是nil
// 好 := hello() ，这是输出的就是nil 的，因为赋值的是hello的返回值
func hello() []string {  
     return nil
 }

 func main() {  
     h := hello
     if h == nil {
         fmt.Println("nil")
     } else {
        fmt.Println("not nil")
    }
}
```

### 15.array
```bazaar
// 数组定义是 [n] n存在
// 不同n代表不同类型数组，不同类型不可以做比较

```