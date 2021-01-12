package write

// 面试题 01. 判定字符是否唯一

// 不要直接上手 就是bitset  ,位运算, ---> 思考,字符到到底是多少个? 那种字符环境? Unicode么?
// 开头 先问 范围
// 字符串的字符范围，如果我告诉他，26个小写英文字母，那可能一开头直接判断如果字符长度>26, 直接返回False
// ascii字符集，然后他的代码里有边界检查，并且针对不同的范围有不同的侧重点，比如说ascii字符集，那也就是128个可能性，16个字节的位运算比较好
// unicode，没有字符范围，老老实实排序再判断是比较符合我对工程师的要求的，因为算法性能稳定，没有额外资源要求，一眼看出没什么不可预见的风险

// 所有第一步判断特殊情况 空时 如何?

// 本题已经给出范围
// 0 <= len(s) <= 100
// 如果你不使用额外的数据结构，会很加分。

// 第一步先写暴力解法 暴力解法
func isUnique(astr string) bool {
	// 这不判断加分
	if len(astr) == 0 {
		return true
	}
	// 暴力求解,第一步加分
	for i := range astr {
		for j := i + 1; j < len(astr); j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}

// 先排序 ,在对比前后字符是否相同 ,不同true ,相同 返回false ,特别大的用快排,其实小点儿的不必优化

// 这里位运算默认是 26字母 ,但是其实不能这么直接做
// func isUnique(astr string) bool {
// 	if len(astr) == 0 {
// 		return true
// 	}
// 	num := 0
// 	for _, v := range astr {
// 		// 相对于 a  的偏移量
// 		moveBit := v - 'a'
// 		// 将1 左移 偏移量个 位 , 1 后面补上 偏移量个 0 ,与 num 相 & 是否为0?
// 		if num&(1<<moveBit) != 0 {
// 			return false
// 		} else {
// 			num = num | (1 << moveBit)
// 		}
// 	}
// 	return true
// }
