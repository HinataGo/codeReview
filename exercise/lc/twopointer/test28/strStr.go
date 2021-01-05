package test28

// 线性时间复杂度 O((N−L)L)  空间 O(1)
//
func strStr(haystack string, needle string) int {
	n := len(haystack)
	n2 := len(needle)
	for i := 0; i < n-n2+1; i++ {
		if haystack[i:i+1] == needle {
			return i
		}
	}
	return -1
}
