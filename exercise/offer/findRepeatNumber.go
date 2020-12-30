package code

// 剑指 Offer 03. 数组中重复的数字
// sort 之后对比数组下标和value 不一致为重复value ,拿数字与下标相同的元素交换

func findRepeatNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
