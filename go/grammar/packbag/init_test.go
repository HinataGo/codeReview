package packbag

/*
import (
	"fmt"
	// 用户自定义包
	"userPackage"
)

func init() {
	fmt.Println("tool init")
}

func main() {
	fmt.Println("main run")
	// 使用userPackage
	// 自定义函数
	userPackage.SomeFunc()
}
*/
// 执行结果  :会先输出 "tool init"，再输出 "main run"
// init & main
// 同 : 1. 两个函数在定义时不能有任何的参数和返回值。
//  2. 该函数只能由 go 程序自动调用，不可以被引用

// init可以重复多个定义,main 不可以,只能一个
// 执行顺序看图 ,先加载init

// 对同一个 go 文件的 init( ) 调用顺序是从上到下的。
//
// 对同一个 package 中的不同文件，将文件名按字符串进行“从小到大”排序，之后顺序调用各文件中的init()函数。
//
// 对于不同的 package，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 init() 函数。
//
// 如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化，
// 例如：导入顺序 main –> A –> B –> C，则初始化顺序为 C –> B –> A –> main，
// 一次执行对应的 init 方法。main 包总是被最后一个初始化，因为它总是依赖别的包

// PS : 避免出现循环 import，例如：A –> B –> C –> A。
// 一个包被其它多个包 import，但只能被初始化一次
