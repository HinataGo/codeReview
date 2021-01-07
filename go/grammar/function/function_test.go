package function

import (
	"fmt"
	"sync"
	"testing"
)

// go函数特点
// 无需前置声明
// 不支持命名嵌套定义 nested
// 不支持同名函数重载 overload
// 不支持默认参数
// 支持不定长参
// 支持多返回值
// 支持匿名函数和闭包
// 太基础的略过,记录一些容易疏忽的,陌生的

// 函数属于 第一类对象,具有相同签名(参数,返回值列表) 的 视作统一类型
// 第一类对象 : 运行期创建,做函数参数或返回值,可存入变量的实体--- 常见的 匿名函数
func other() {
	println("test206")
}

// 函数可以作为参数传递
func this(f func()) {
	// 执行参数 传来的的 f 函数
	f()
}

func TestDefine(t *testing.T) {
	f := other
	this(f) // 成功打印数据
}

// 推荐使用命名类型定义函数
// ps:  ... 为可变参数  ...interface{} ,fun期望类型的第一个参数string，然后是类型的0到N个参数interface{}。
// 表示函数可以接收可变数量的参数，通常称为varargs（或var-args或其他拼写形式）。这称为可变函数。
// 尽管提供任何数量的参数的功能似乎非常方便，并且在这里不必过多讨论细节，否则可能会出现偏离主题的风险，但根据语言的实现，它会带来一些警告：
// 增加内存消耗，
// 可读性下降，
// 降低代码安全性。
// An interface (or closer to go's paradigm, a protocol), is a type that defines a contract for other objects to comply to
type fun func(string, ...interface{}) (string, error)

// 非命名则会这样
func fun2(f1 fun, string2 string, a ...interface{}) (string, error) {
	return f1(string2, a)
}

// 函数智能判断是否为nil ,不可以比较两个函数
// 函数返回局部指针是安全的,编译器会通过逃逸分析,来决定是否在堆上分配内存

// 变量遮蔽,命名返回值
func anonymous(t *testing.T) (i int, e error) {
	i = 1
	// 这里变量遮蔽,但同时需要命名所有return 值,否则编译器无法识别,golang会自己跳过未命名返回值
	return
}

// 匿名函数
// 匿名函数可以做参数,返回值,赋值给变量
func TestAnonymous(t *testing.T) {
	// 直接执行
	func(string2 string) {
		println(string2)
	}("hello")

}

// 另一种写法
func TestFun3(t *testing.T) {
	n := 3
	a := func(s int) int {
		n *= s
		return n
	}
	// 这里 a是函数 后面调用得加上 小括号 (参数可选)
	fmt.Println(a(1))
}

// TODO 闭包通过求值引用,可能导致生命周期延长,甚至被分配到堆内存

// 闭包导致  延迟求值( )
// TestBag 输出结果
// 0xc000094230 3
// 0xc000094230 3
// 0xc000094230 3
func getN() []func() {
	var s []func()
	for i := 0; i < 3; i++ {
		s = append(s, func() {
			println(&i, i)
		})
	}
	return s
}
func TestBag(t *testing.T) {
	for _, f := range getN() {
		f()
	}
}

// TODO defer延迟调用,很简单,暂时不写

// 性能 defer较差(注册,调用,额外缓存开销等等)  所以可以直接电泳CAll 汇编指令调用函数
// 输出结果(go1.15 ) 明显 效果相差 call更有, 旧版本测试相加几倍,, 1.15测试效果还行,应当减少使用defer
// BenchmarkCall
// BenchmarkCall-12         	73136824	        15.8 ns/op
// BenchmarkDeferCall
// BenchmarkDeferCall-12    	66714849	        18.0 ns/op
var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	defer m.Unlock()
}

// testing.B使用 Benchmark // T使用TestXxx
func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkDeferCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}

// error处理 TODO
// panic recover 类似以 try catch 结构化异常
// 它们是内置函数,不是语句
// panic会立即中断当前的函数流程,执行延迟调用, 且是空接口类型,支持任何对象做错误返回状态
// 延迟调用中,recover可以捕获并返回panic 提交错误的对象,
// 连续调用panic 仅最后一个会被recover捕获
// 延迟函数中panic,不会影响后续延迟调用执行, recover之后的panic可以再次被捕获
// TODO ps 除非错误不可恢复,系统无法正常工作,否则别用panic(文件系统无权限,端口占用,数据库未启动等等)
