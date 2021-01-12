package codeoffer

// [[1, 1, 1, 1, 1, 1, 1],
// [1, 2, 2, 2, 2, 2, 1],
// [1, 2, 3, 3, 3, 2, 1],
// [1, 2, 2, 2, 2, 2, 1],
// [1, 1, 1, 1, 1, 1, 1]]
//
// 顺时针输出矩阵
func spiralOrder(matrix [][]int) []int {
	// 这么写判断会在 matrix 时 [] 不存在 matrix[0] 报错
	// row, columns := len(matrix), len(matrix[0])
	// if row == 0 || columns == 0 {
	// 	return []int{}
	// }
	if len(matrix) == 0 {
		return []int{}
	}
	row, columns := len(matrix), len(matrix[0])
	res := make([]int, row*columns)
	index := 0
	// l左边界 , r右边界 , t 上边界, b底部,下边界
	// 顺时针思路,逆时针则相反
	// 到达右边界就 -- > 向下遍历
	// 到达下边界就 -- > 向左遍历
	// 到达左边界就 -- > 向上遍历
	// 到达上边界就 -- > 向右遍历
	l, r, t, b := 0, columns-1, 0, row-1
	for l <= r && t <= b {
		// 第一遍 右遍历
		for i := l; i <= r; i++ {
			res[index] = matrix[t][i]
			index++
		}
		// 向下遍历,先下一一层 赋值给j,遍历 直到矩阵底端
		for j := t + 1; j <= b; j++ {
			res[index] = matrix[j][r]
			index++
		}
		// 此时如果 l未达到r (左右不重合) 并且 底端和顶端为重合
		// 检测是否遍历完
		if l < r && t < b {
			// i 左移一位,开始向左遍历.要大于l
			for i := r - 1; i > l; i-- {
				res[index] = matrix[b][i]
				index++
			}
			// 到大最左边,向上遍历
			for j := b; j > t; j-- {
				res[index] = matrix[j][l]
				index++
			}
		}
		// 一轮结束后需要缩小最外围一圈
		// 左边界右移 ,右边界左移, 上边界下移, 下边界上移
		l++
		r--
		t++
		b--
	}
	return res
}
