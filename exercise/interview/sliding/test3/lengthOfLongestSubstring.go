package test3

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	m := make(map[byte]int, 0)

	r, res := -1, 0
	for l := range s {
		if l != 0 {
			delete(m, s[l-1])
		}
		for r+1 < n && m[s[r+1]] == 0 {
			m[s[r+1]]++
			// 右指针右移
			r++
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
