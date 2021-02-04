package lc

import "math"

var pre = math.MinInt64

func isValidBST(root *TreeNode) bool {
	// 再次初始化,防止数据被修改后出现问题
	pre = math.MinInt64
	return is(root)
}

func is(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !is(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return is(root.Right)
}
