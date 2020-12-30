package code

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if math.Abs(float64(l-r)) < 2 {
		return int(math.Max(float64(l), float64(r)) + 1)
	} else {
		return -1
	}
}
