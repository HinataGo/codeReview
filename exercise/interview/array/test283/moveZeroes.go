package test283

// 双指针,移动0 ,l,r
// 一快一慢,有伴非零和左边交换,一步步交换,数组原地操作,tips.不要投机取巧,最后几位直接置零是不合题意的
func moveZeroes(nums []int) {
	n := len(nums)
	l, r := 0, 0
	for r < n {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
