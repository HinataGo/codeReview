package test20

// 本题目归类为栈, 很简单  有效的括号(
//    左括号必须用相同类型的右括号闭合。
//    左括号必须以正确的顺序闭合。)
//	一个个往栈里存,遇到对应点额括号 ,pop一个 (弹出栈首) 输出最后的结果判断是否等于 "" 即可
func isValid(s string) bool {
	n := len(s)
	// 奇数个明显不匹配,直接出结果
	if n%2 == 1 {
		return false
	}
	// 首先创建一个,table 存储对应的数据
	// 这里我们是看右括号匹配左括号 一一删
	table := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	// 思想就是找到了 右括号,再看看它前面那个是不是左括号,因为得对称才能符合要求
	for i := 0; i < n; i++ {
		// 查询 目前扫到的字符是否匹配table 中的右括号
		if _, ok := table[s[i]]; ok {
			// 这是时找打了右括号 ,然后 第二部分条件--匹配 左右括号是否符合值
			if len(stack) == 0 || stack[len(stack)-1] != table[s[i]] {
				return false
			}
			// pop操作 ,就是删除栈首
			stack = stack[:len(stack)-1]
		} else {
			// 左括号直接添加进入 stack
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
