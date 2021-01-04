package test13

// 转义罗马字符为整数
// 换个思路,一个个数据读取没问题, 左边的值要是比自己小那么当前结果就是减去左边的值,否否则直接加上当前值即可
// 下面是从0 -n遍历,换个思路就是从 n-0 遍历可以优化一些
// func romanToInt(s string) int {
// 	tmp := []byte(s)
// 	res := 0
// 	for i := 0; i < len(tmp); i++ {
// 		if  i - 1 < 0 {
// 			res = getNum(tmp[i])
// 			continue
// 		}
// 		pre := getNum(tmp[i-1])
// 		now := getNum(tmp[i])
// 		if now > pre {
// 			res = res + now - pre*2
// 		}else{
// 			res += now
// 		}
// 	}
// 	return res
// }

// 较为优雅的一种写法
func romanToInt(s string) int {
	r, res := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		now := getNum(s[i])
		// 左大右小,因此直接加就可以
		if now >= r {
			res += now
		} else {
			// 符合 左小右大情况, 直接 减去当前的值 ( 5 - 1 = 4 ,这里的5 是之前已经加上了)
			res -= now
		}
		// 最后更新右边一个数的值
		r = now
	}
	return res
}

func getNum(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
