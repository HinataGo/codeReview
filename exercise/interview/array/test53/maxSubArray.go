package test53

// 最大子序和
// dp
// 从0 遍历 ,依次当前一个数 ,对比大小,增加就收了,减少就抛弃
// ps: 这里只要求结果,不要求
func maxSubArray(nums []int) int {
	n := len(nums)
	res := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		// 取较大的结果进行保留
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
