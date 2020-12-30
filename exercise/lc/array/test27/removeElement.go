package test27

// 大家可以参考283题官方解答,一模一样的
// 并非原创,仅供学习
func removeElement(nums []int, val int) int {
	n := len(nums)
	l, r := 0, 0
	// 利用右指针遍历(较快的那个)
	for r < n {
		// 当nums的值不是目标删除的val时
		if nums[r] != val {
			// 交换元素
			nums[l], nums[r] = nums[r], nums[l]
			// 向右一定左指针
			l++
		}
		// 保存右指针向右移动直到达到数组末尾,并且始终比做指针l快
		r++
	}
	// 新长度就是已经删除的部分长度,因为只考虑到l
	return l
}
