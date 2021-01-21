package test242

// 排序思路
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 很简单，字符串转 字符数组 ，内部排序， 转字符串对比
// func isAnagram(s string, t string) bool {
// 	s1 := []byte(s)
// 	t1 := []byte(t)
// 	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
// 	sort.Slice(t1, func(i, j int) bool { return t1[i] < t1[j] })
// 	return string(s1) == string(t1)
// }

// hash思路
// 1.确定26字母时的做法
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	var a1, a2 [26]rune
// 	for i := range s {
// 		a1[s[i-'a']]++
// 	}
// 	for i := range t {
// 		a2[t[i-'a']]++
// 	}
// 	return a1 == a2
// }

// 不确定字符范围，假设咋Unicode范围内，就不可以用上面的范围
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	res := make(map[rune]int, 0)
	for _, v := range s {
		if res[v] != 0 {
			res[v]++
		}
	}
	for _, v := range t {
		if res[v] != 0 {
			res[v]--
		} else {
			return false
		}
	}
	return true
}
