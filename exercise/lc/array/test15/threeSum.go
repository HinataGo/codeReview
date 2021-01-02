package test15

import "sort"

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。

// 排序 之后按照普通 两数之和做
func threeSum(nums []int) [][]int {
	n := len(nums)
	// go自带快速排序
	sort.Ints(nums)
	var res [][]int
	// 边界确认
	if n < 3 {
		return res
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		r := n - 1
		// 目标值的相反数
		target := -1 * nums[i]
		// 这个过程结合下高中数学常见的思想,分情况讨论问题,
		// 我们所有的遍历和优化取舍,都是在一定条件下
		// 所以,这里 我们先思考划分条件,在进行code ,记得思路和code分开
		// 解题思路优先,code大部分还是对语法的熟练和理解
		// 这里其实可以视为三指针, i 为主,头到尾遍历
		// l, r都是i 后面的元素, 然后这样保证i 始终变化,不重复,是的不会有相同的三元组(三个元素的数组)
		// l从i的右边开头,r则是从末尾.然后寻找,这很容易理解
		// 当 nums[i] 的值大于0 时,可以发现,右边都是大于零的,这里可以提前终止,英文三个大于0 的数之和不会是0
		// 记住,这里的遍历时在l 的基础上
		for l := i + 1; l < n; l++ {
			// 不断的寻找和推进 i 过程,
			// 因为 l 和 i 的值相同并且,不符合预期,因此跳过当前 l
			// 这里不要写 l >= i +1,因为后边跳跃这一步其实是,
			// 检测 num[i+1] == nums[i] ,并且三数之和不符合要求
			if l > i+1 && nums[l] == nums[l-1] {
				continue
			}
			// 两数之和过大,怎么办? ,肯定减少那个大数字 r 看看能不能满足要求
			// 右指针想左 移动
			for l < r && nums[l]+nums[r] > target {
				r--
			}
			// l 和 r相遇,说明不必遍历了,可以退出双指针遍历
			// 这里还要思考为什么 l== r 要放在这儿?
			if l == r {
				break
			}
			// 当寻找到目标数据时,添加我们需要的值到 结果res 数组中去
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
			}
		}
	}
	// 我们返回 结果res
	return res
}
