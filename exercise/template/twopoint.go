package template

// 将 普通的双层遍历优化为 O(n^2) O(n)

func twoPoint(n int) {
	res := 0
	for i, j := 0, 0; i < n; i++ {
		for j < i && check(i, j) {
			j++
		}
		res = max(res, i-j+1)
	}
}

func max(res int, i int) int {
	if res > i {
		return res
	}
	return i
}

func check(i, j int) bool {
	// TODO 条件
	return true
}
