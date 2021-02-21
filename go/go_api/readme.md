# go 各种api操作学习
- go中的关键字 25 个,但是像bool、byte、error、true、iota ,int 
  在Go中它们被称为 预定义标识符 --(builtin/builtin.go) 在这个包中可以找到
- 内建函数仅仅是一个标识符，在Go源码编译期间，
  Go编译器遇到内建函数标识符时会将其替换为若干runtime的调用
#### 关键字
```go
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
```
#### 预定义标识符
``go
内建常量: true false iota nil

内建类型: int int8 int16 int32 int64
          uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error

内建函数: make len cap new append copy close delete
          complex real imag
          panic recover
```