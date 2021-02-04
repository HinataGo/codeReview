package lc

func maxDepth(root *TreeNode) int {
	ret := 0
	if root == nil {
		return 0
	}
	ret++
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return ret + left
	}
	return ret + right
}
