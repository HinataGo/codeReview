package test27

import "sort"

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, 0
	for r < n {
		if nums[r] != val {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
	return l
}
