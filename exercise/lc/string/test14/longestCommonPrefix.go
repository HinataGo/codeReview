package test14

import "sort"

func longestCommonPrefix(s []string) string {
	// 因为是找前缀,二话不说,先拿一个字符的开头和后面一个对比,重复部分提取出来ret
	// ret 再次和后面第三个的比较.够长就保留.不够就删除,如果过程中恰好删完则结束,否则ret 返回即可
	// 可以先排序,ps: 字符串排序 ,首字母也是顺序, 所以是这个特点,因此这样比较首部是否合理即可
	res := ""
	n := len(s)
	if n == 0 {
		return ""
	}
	sort.Strings(s)
	l, r := s[0], s[n-1]
	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			break
		}
		res = l[:i+1]
	}

	return res
}
