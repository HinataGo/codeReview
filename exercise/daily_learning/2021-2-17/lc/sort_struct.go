package lc

import (
	"fmt"
	"sort"
)

// 对结构体指定字段进行排序
type User struct {
	Name string `json:"name"` // `json:"xxx"`：在结构体和json字符串字段顺序不一致的情况下：unmarshal根据tag去寻找对应字段的内容
	Age  int    `json:"age"`
}

// type Users []User
// func SortByAge(u Users) {
func SortByAge(u []User) {
	fmt.Printf("源数据：%+v\n", u)

	sort.Slice(u, func(i, j int) bool { // desc
		return u[i].Age > u[j].Age
	})
	fmt.Printf("按Age降序：%+v\n", u)

	sort.Slice(u, func(i, j int) bool { // asc
		return u[i].Age < u[j].Age
	})
	fmt.Printf("按Age升序：%+v\n", u)
}

func main() {
	// 初始化结构体对象数组：
	// 初始化方法1：
	// users := Users{
	// 	{
	// 		Name: "test1",
	// 		Age:  22,
	// 	},
	// 	{
	// 		Name: "test2",
	// 		Age:  19,
	// 	},
	// 	{
	// 		Name: "test3",
	// 		Age:  25,
	// 	},
	// }

	// 初始化方法2：
	var users []User
	var u User
	u.Name = "test1"
	u.Age = 22
	users = append(users, u)
	u.Name = "test2"
	u.Age = 20
	users = append(users, u)
	u.Name = "test3"
	u.Age = 26
	users = append(users, u)

	SortByAge(users)
}
