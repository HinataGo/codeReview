package test28

// 滚动哈希：常数时间生成哈希码
// Rabin Karp - 常数复杂度
// 这才是最好的解决方法
// 好了写一半 TODO
func strStr3(haystack string, needle string) int {
	h, n := len(haystack), len(needle)
	if n > h {
		return -1
	}
	// TODO 没想好
	return -1
}

func charToInt(idx int, s string) int {
	return int(s[idx]) - int('a')
}
