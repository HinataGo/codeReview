package test20

func isValid(s string) bool {
	// 常见解法，用一个stack 去匹配数据，pop 和 push
	// string 的每个字符单位是 rune
	n := len(s)
	if n%2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	// 建立字符信息匹配表
	table := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		// 查找hash表，查不到就是0  查到非0
		if table[v] != 0 {
			// 两者其一不符合不符合才算错误，栈空，
			// 并且 栈 的top 元素必须是和 左括号 v 对应的右括号，再table得查表查到
			if len(stack) == 0 || stack[len(stack)-1] != table[v] {
				return false
			}
			// 删除top元素就是重新赋值一个新切片，由于片取值 左开右闭 ，所以删除top是这样操作的
			stack = stack[:len(stack)-1]
		} else {
			// 查到的v是左括号，添加进入stack
			stack = append(stack, v)
		}
	}
	// 最后stack 全空才是正确的，需要检查 这是为了防止 （（（（ 这种类似情况的出现
	return len(stack) == 0
}
