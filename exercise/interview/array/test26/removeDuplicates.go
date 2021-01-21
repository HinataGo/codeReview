package test26

func removeDuplicates(nums []int) int {
	// 数组已经排序了
	// 思路一，比较 现在和之前 ， 相同将当前换成后面一个，但是这样复杂度太高，而且后面还有很多难操作的
	// 这里使用双指正优化
	// 左指针 固定， 有指针移动直到两个数不同，
	// 有指针和 左指针右边一个交换数据
	// 将左指针右移动一位
	l, r := 0, 0
	for r < len(nums) {
		if nums[l] != nums[r] {
			nums[l+1], nums[r] = nums[r], nums[l+1]
			l++
		}
		r++
	}
	return l + 1
}
