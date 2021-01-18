package test11

// 盛水最多的容器
// 双指针
// 左右指针同时向内缩减，随后 找准什么时候缩小那边即可，很easy 的想法就是，谁短移动谁，
// 后续优化还可以看，下一个柱子变高就可以要，否则往后跳，这样优化后省略许多无用的计算，最后得出max 即可
func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	if n <= 1 {
		return 0
	}
	res := (r - l) * min(height[l], height[r])
	for l < r {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		tmp := (r - l) * min(height[l], height[r])
		res = max(tmp, res)
	}

	return res
}

func max(tmp int, res int) int {
	if tmp > res {
		return tmp
	}
	return res
}

func min(i int, i2 int) int {
	if i > i2 {
		return i2
	}
	return i
}
