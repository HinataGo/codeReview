package typesof

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// 自定义类型 type  很类似 define
// type支持组定义,
type (
	flags byte
	user struct {
		name string
		age uint8
	}
	event func(string) bool				// 函数类型
)

// 未命名类型 数组,切片,字典,通道等类型与元素类型和长度等有关---- 可以使用type 改变命名类型
// 可以视作统一类型
// 相同  基本类型的指针
// 元素类型和长度的数组
// 元素类型的切片
// 相同键值类型的字典
// 相同数据类型及操作方向的通道
// 相同字段的结构体
// 相同签名的函数(参数和返回值列表,不包含参数名)
// 相同方法集(方法名,方法签名,不包括顺序)的接口

// struct tag (tag也是类型的一部分额,不仅仅是元数据描述)
// ``(反引号)：反引号用来创建 原生的字符串字面量,这些字符串可能由多行组成(不支持任何转义序列)，
// 原生的字符串字面量多用于书写多行消息、HTML以及正则表达式。
/*
在golang中，命名都是推荐都是用驼峰方式，并且在首字母大小写有特殊的语法含义：包外无法引用。
但是由经常需要和其它的系统进行数据交互，例如转成json格式，存储到mongodb啊等等
。这个时候如果用属性名来作为键值可能不一定会符合项目要求。
*/
func TestTag(test *testing.T) {
	type User struct {
		UserId   int    `json:"user_id" bson:"user_id"`
		UserName string `json:"user_name" bson:"user_name"`
	}
	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
	// 输出内容：{"user_id":1,"user_name":"tony"}
	// 如果在属性中不增加标签说明，则输出：{"UserId":1,"UserName":"tony"}
	// 可以看到直接用struct的属性名做键值。
	// 其中还有一个bson的声明，这个是用在将数据存储到mongodb使用的。

	// 获取tag中的内容
	s := reflect.TypeOf(u)
	field := s.Elem().Field(0)
	fmt.Println(field.Tag.Get("json")) // 输出：user_id
	fmt.Println(field.Tag.Get("bson")) // 输出：user_id
	s1 := reflect.TypeOf(u).Elem() // 通过反射获取type定义
	for i := 0; i < s1.NumField(); i++ {
		fmt.Println(s1.Field(i).Tag) 		// 将tag输出出来
	}
}
