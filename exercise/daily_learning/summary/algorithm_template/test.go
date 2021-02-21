package algorithm_template

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 如果找不到，left 是第一个大于target的索引
	// 如果在B+树结构里面二分搜索，可以return left
	// 这样可以继续向子节点搜索，如：node:=node.Children[left]
	return -1
}
