package test15

import "sort"

// 更加符合go的一种写法
func threeSum2(nums []int) [][]int {
	// 按照从小到大排序
	sort.Ints(nums)
	width := len(nums)
	var res [][]int

	// 边界确认
	if width < 3 {
		return res
	}
	// [0,0,0,0]
	// 固定一个值，用首位指针进行遍历相加
	for i, value := range nums {
		// 因为前面已经排序了，所以当前值为正数时，后面相加肯定不可能为0，所以直接跳出
		if value > 0 {
			return res
		}
		// 如果索引大于0，并且前后两个值相同的话，可以跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 定义首尾指针
		l := i + 1
		r := len(nums) - 1
		for l < r {
			// 判断指针的值
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
