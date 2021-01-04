package test14

// 写个通俗解法,没有啥奇淫技巧,就是一个个遍历 ,最简单的
// 因为是找前缀,二话不说,先拿一个字符的开头和后面一个对比,重复部分提取出来ret
// res 再次和后面第三个的比较.够长就保留.不够就删除,如果过程中恰好删完则结束,否则res 返回即可
// 这里解法的关键在于 拿出res 作为一个基准 ,然后r一遍扫描过 整个数组, 此时数组不需要像之前那种解法 先排序
// 但是前一种解法也有可取之处,只比较两个数,因此也是非常高效的(这里排序用快排)
func longestCommonPrefix2(s []string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	res := s[0]
	for i := 1; i < n; i++ {
		r := s[i]
		j := 0
		length := getMin(res, r)
		// 当符合要求时 继续下一个字符的判断
		for j < length && res[j] == r[j] {
			j++
		}
		// 更新res
		res = res[:j]
		if len(res) == 0 {
			break
		}
	}
	return res
}

func getMin(l, r string) int {
	if len(l) > len(r) {
		return len(r)
	}
	return len(l)
}
