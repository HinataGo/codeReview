package test9

// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 121 true  -121 false

// 目前知道的最优 ,
func isPalindrome(x int) bool {
	res := 0
	// 负数 ,末尾为0 的都不是正解
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 之前的想法是解出反转的数在判断 ,后来发觉其实,不用全算完 ,因为回文数会对称,因此计算一半即可
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	// 第二处理 121 这种 res获得的是1 不是12 这种情况
	return x == res || x == res/10
}
