package initialization

type ByteSize float64

// 它们是在编译时创建的，即使在函数中定义为局部变量时也只能是数字，字符（符文），字符串或布尔值。

// 由于编译时的限制，定义它们的表达式必须是可由编译器评估的常量表达式
// 1 << 3 , math.Sin(math.Pi/4) 都不可以被定义程常量, 因为需要计算,他们的值实在运行时出来的
const (
	_           = iota // 通过分配给空白标识符来忽略第一个值 0
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
