package code

func replaceSpace(s string) string {
	// 1. 不清楚怎么在数组里添加过长的字符 原先字符只有一个,现在是两个(主要是原地做是不是太复杂了)
	// 题解参考 ,包数据拷贝到另外一个数组里吧

	var res string
	for i, v := range s {
		if v == ' ' {
			res += "%20"
		} else {
			res += string(s[i])
		}
	}
	return res
}
