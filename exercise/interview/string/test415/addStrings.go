package test415

import "strconv"

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

//    num1 和num2 的长度都小于 5100
//    num1 和num2 都只包含数字 0-9
//    num1 和num2 都不包含任何前导零
//    你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式

func addStrings(num1, num2 string) string {
	add := 0
	res := ""
	for n1, n2 := len(num1)-1, len(num2)-1; n1 >= 0 || n2 >= 0 || add != 0; n1, n2 = n1-1, n2-1 {
		var x, y int
		if n1 >= 0 {
			x = int(num1[n1] - '0')
		}
		if n2 >= 0 {
			y = int(num2[n2] - '0')
		}
		result := x + y + add
		// strconv.Itoa 它可以将数字转换成对应的字符串类型的数字
		res = strconv.Itoa(result%10) + res
		add = result / 10
	}
	return res
}
