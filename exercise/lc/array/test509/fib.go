package test509

// 过于简单 ,由于题目 0 <= n <= 30
func fib(n int) int {
	a, b, res := 0, 1, 0
	for i := 0; i < n; i++ {
		a = b
		b = res
		res = a + b
	}
	return res
}
