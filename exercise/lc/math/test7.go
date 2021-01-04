package math

import "math"

// 为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0
// 思路很easy 就是 从各位上开始拿数字,每次去一个各位数字,然后对他*10 ,同时 *10 之前判断res 是否溢出?
// 最终结果即为正确的翻转数字
func reverse(x int) int {

	res := 0
	for x != 0 {
		// 每次向前减少一个10进制值位 并且取出各位上的值
		pop := x % 10
		x /= 10
		// 溢出
		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && pop > 7 {
			return 0
		}
		// 溢出
		if res < math.MinInt32/10 || res == math.MinInt32/10 && pop < -8 {
			return 0
		}
		res = res*10 + pop
	}
	return res
}
