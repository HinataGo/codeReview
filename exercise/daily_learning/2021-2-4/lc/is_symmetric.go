package lc

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(l, r *TreeNode) bool {
	if l == r {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	// l,r 都非空时 去return 那一步判断 val 值
	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}
