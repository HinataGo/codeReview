package test66

// 不是很优雅吗，但是思路很清晰
func plusOne(digits []int) []int {
	n := len(digits)
	cout := 1
	res := make([]int, 0)
	i := n - 1
	for i >= 0 {
		if digits[i]+cout > 9 {
			digits[i] = 0
			cout = 1
		} else {
			digits[i] += cout
			cout = 0
			break
		}
		i--
	}
	if i == -1 {
		res = append(res, 1)
		res = append(res, digits...)
	} else {
		res = digits
	}
	return res
}
