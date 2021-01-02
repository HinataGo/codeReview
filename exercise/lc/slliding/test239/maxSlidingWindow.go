package test239

// 暴力解法 --- 超时 ,转解法2
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	tmp := make([]int, k)
	res := make([]int, n-k+1)
	if k == 1 {
		return nums
	}
	for i := 0; i <= n-k; i++ {
		for j := i; j < i+k && i+k <= n; j++ {
			tmp[j-i] = nums[j]
		}
		res[i] = getMax(tmp)
	}
	return res
}
func getMax(arr []int) int {
	tmp := arr[0]
	for _, i2 := range arr {
		if i2 > tmp {
			tmp = i2
		}
		continue
	}
	return tmp
}
