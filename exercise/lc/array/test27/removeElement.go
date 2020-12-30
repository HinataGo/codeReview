package test27

// 同238 移动0 ,一样的题
func removeElement(nums []int, val int) int {
	n := len(nums)
	l, r := 0, 0
	for r < n {
		if nums[r] != val {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
	return l
}
