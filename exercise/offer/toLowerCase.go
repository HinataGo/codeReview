package code

// 709. 转换成小写字母
// ASCII 码转换即可
func toLowerCase(str string) string {
	res := []byte(str)
	for i := range res {
		if res[i] >= 'A' && res[i] <= 'Z' {
			res[i] = res[i] + 32
		}
	}
	return string(res)
}
