package lc

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	l, r, sum, res := 0, -1, 0, n+1
	for l < n {
		if r+1 < n && sum < s {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			res = min(res, r-l+1)
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}
