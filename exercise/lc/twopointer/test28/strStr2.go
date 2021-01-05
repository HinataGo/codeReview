package test28

// 双指针
// 最坏时间复杂度为 O((N−L)L)，最优时间复杂度为 O(N)
// 空间O(1)
// 优化写法
func strStr2(haystack string, needle string) int {
	h, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	// 注意这里写 h -n +1 (比如 两个都为 "a" 时 ,不加1 会导致直接进不去循环)
	for i := 0; i < h-n+1; i++ {
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == n-1 {
				return i
			}
		}
	}
	return -1
}

// 最初的想法
// func strStr2(haystack string, needle string) int {
// 	// 取出两个字符长度
// 	n1, n2 := len(haystack),len(needle)
// 	// 题意,needle 长度为0 返回0
// 	if n2 == 0 {
// 		return 0
// 	}
// 	p := 0
// 	// 遍历 两个指针 p指针用来扫描数据
// 	// cur指针用于计数
// 	// p指针在 haystack中 扫描不可能超过 n1- n2 + 1 因为要匹配needle字符
//
// 	for p < n1-n2+1 {
// 		// 首字母不相同 ,右移 p指针 在haystack 中继续扫描,直到 首字母相同或者 ,p超出限制
// 		for p < n1-n2+1 && haystack[p] != needle[0] {
// 			p ++
// 		}
// 		// 提前终止条件
// 		if p == n1 - n2 + 1{
// 			return -1
// 		}
// 		// 双指针定义
// 		l , r := 0, 0
// 		// 匹配 字符相同 所有指针右移
// 		for r < n2 && p < n1 && haystack[p] == needle[r]{
// 			p++
// 			r++
// 			l++
// 		}
// 		// l 扫描出的 长度与 n2 一致 ,这里就满足结果了
// 		if l == n2  {
// 			return p - n2
// 		}
// 		p = p - l + 1
// 	}
// 	// 扫描完都不符合,返回 -1
// 	return -1
// }
