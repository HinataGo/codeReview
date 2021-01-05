package test35

func searchInsert2(nums []int, target int) int {
	index := 0
	i := 0
	for i < len(nums) {
		if nums[i] >= target {
			index = i
			return index
		} else {
			i++
		}
	}
	// nums 轮询完都没有插入，则插入末尾
	if i == len(nums) {
		index = len(nums)
	}
	return index
}
