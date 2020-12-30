package main

// 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，
// 如果找到了，直接返回 2 个数字的下标即可。
// 如果找不到，就把这个数字存入 map 中，
// 等待扫到“另一半”数字的时候，再取出来返回结果。

// hash 查找近似 O(1)
// 时间复杂度：O(N)，其中 N 是数组中的元素数量。对于每一个元素 x，我们可以 O(1) 地寻找 target - x。
// 空间复杂度：O(N)，其中 N 是数组中的元素数量。主要为哈希表的开销

// 两种写法 熟悉下语法，思路一致

func twoSum(nums []int, target int) []int {
	// 第一部分可有可无
	// if len(nums) <= 1 {
	//	return []int{0, 0}
	//  }
	table := make(map[int]int)
	// table := map[int]int{}
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if k, v := table[other]; v {
			return []int{k, i}
		}
		table[i] = i
	}
	// for i,x := range nums{
	//	if k, v := table[target - x]; v {
	//		return []int{k, i}
	//	}
	//	table[x] = i
	// }

	return nil
}
