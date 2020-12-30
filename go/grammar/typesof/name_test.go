package typesof

import (
	"strconv"
	"testing"
)

// 命名.遵循英文命名,驼峰式即可
// _ 空标识符,用于忽略左值,无法读取,临时规避编译器对未使用变量/导入包的 错误检查,它是预置成员,不可重新定义
// 预置成员
func TestName(t *testing.T) {
	i , _ := strconv.Atoi("1")
	println(i)
}