package test15

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for l := 0; l < n && nums[l] <= 0; l++ {
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}

		for i, r := l+1, n-1; i < r; i++ {
			tmp := nums[i] + nums[l] + nums[r]
			if i > l+1 && nums[i] == nums[i-1] {
				continue
			}
			for tmp > 0 {
				r--
				// 防止溢出
				if i == r {
					break
				}
				tmp = nums[i] + nums[l] + nums[r]
			}
			if tmp == 0 && i < r {
				res = append(res, []int{nums[l], nums[i], nums[r]})
			}
		}
	}
	return res
}
