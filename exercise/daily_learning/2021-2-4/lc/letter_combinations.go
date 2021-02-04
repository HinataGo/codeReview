package lc

// 17. 电话号码的字母组合
/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

*/
var phoneMap = []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
var res []string

func find(digits string, index int, s string) {
	if index == len(digits) {
		// 保存digits
		res = append(res, s)
		return
	}
	v := digits[index]
	letters := phoneMap[v-'0']
	for i := 0; i < len(letters); i++ {
		find(digits, index+1, s+string(letters[i]))
	}
}
func letterCombinations(digits string) []string {
	// 每次初始化全局变量
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	find(digits, 0, "")
	return res
}
