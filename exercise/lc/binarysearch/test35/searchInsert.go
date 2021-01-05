package test35

// 很暴力
func searchInsert(nums []int, target int) int {
	n := len(nums)
	res := -1
	if nums[0] >= target {
		return 0
	}
	if nums[n-1] < target {
		return n
	}
	for i := 0; i < n-1; i++ {
		if nums[i] < target && nums[i+1] >= target {
			res = i + 1
		}
		if nums[i] == target {
			res = i
		}
	}
	return res
}
