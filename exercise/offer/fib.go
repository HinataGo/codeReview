package code

// 剑指 Offer 10- I. 斐波那契数列
func fib(n int) int {
	a, b, res := 0, 1, 0
	for i := 0; i < n; i++ {
		res = (a + b) % 1000000007
		a = b
		b = res
	}
	return a
}
