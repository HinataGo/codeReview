package test70

// 注意：给定 n 是一个正整数。
// dp思路
func climbStairs(n int) int {
	// 递推公式 fn = f(n-1) + f(n-2)   n > 0
	a, b, res := 0, 0, 1
	for i := 1; i <= n; i++ {
		a = b
		b = res
		res = a + b
	}
	return res
}
