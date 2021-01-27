package packbag

// 现在使用 go mod
/*
// 单个导入
import "package"
// 批量导入
import (
	"log"
	"hash"
	起别名
	// p1 "package1"
)
*/
// 使用时：别名操作，调用包函数时前缀变成了我们的前缀
// p1.Method()

/*
Go 语言的入口 main() 函数所在的包（package）叫 main，
main 包想要引用别的代码，需要import导入！

src 目录是以代码包的形式组织并保存 Go 源码文件的。
每个代码包都和 src 目录下的文件夹一一对应。每个子目录都是一个代码包。

代码包包名和文件目录名，不要求一致。比如文件目录叫 hello，
但是代码包包名可以声明为 “main”，
但是同一个目录下的源码文件第一行声明的所属包，必须一致！
*/
// 可以被导出的函数
func FuncPublic() {
}

// 不可以被导出的函数
func funcPrivate() {
}
