package test26

// 删除排序数组中的重复项
// 给定数组 nums = [1,1,2],
//
// 函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
// 你不需要考虑数组中超出新长度后面的元素。
func removeDuplicates(nums []int) int {
	n := len(nums)
	m := nums
	j := 0
	for i := 1; i < n; i++ {
		if m[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
