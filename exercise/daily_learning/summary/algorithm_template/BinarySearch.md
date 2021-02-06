# 二分搜素
- 给一个有序数组和目标值，找第一次/最后一次/任何一次出现的索引，如果没有出现返回-1
- 时间复杂度 O(logn)，使用场景一般是有序数组的查找
- ps
    - 1、初始化：start=0、end=len-1 
    - 2、循环退出条件：start + 1 < end
    -  3、比较中点和目标值：A[mid] ==、 <、> target
    -  4、判断最后两个元素是否符合：A[start]、A[end] ?= target
### 二分搜索模板
```bazaar
// 适用 : 查找第一次和最后一次的位置, 包括查找重复元素
// 二分搜索最常用模板 
func search(nums []int, target int) int {
	// 1、初始化left、right
	left, right := 0, len(nums) - 1
	// 2、处理for循环
	for left+1 < right {
		// 小心bug,这里防止越界处理
		mid := left + (right-left)/2
		// 3、比较a[mid]和target值
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		}
	}
	// 4、最后剩下两个元素，手动判断
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}
```
```bazaar
// 精简版本, 不需要第一个和最后一个位置, 或者没有重复元素
// 无重复元素搜索时，更方便
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
```