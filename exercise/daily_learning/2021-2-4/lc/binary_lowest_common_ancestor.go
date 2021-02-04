package lc

// 二叉树的最近公共祖先
// 考虑 pq 是否可以为空
// 这里允许pq 为空第一步测试
func binaryLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 判断p q 在左孩子 还是右孩子
	// 根据二叉搜索树的性质
	// 这里默认 p < q 实际的话得问
	// 如果需要写个swap
	if root == nil {
		return nil
	}
	// 左子树,,那么递归深入去搜索左子树的情况
	if p.Val < root.Val && q.Val < root.Val {
		return binaryLowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return binaryLowestCommonAncestor(root.Right, p, q)
	}
	// 刚好一左一右,那么直接返回root
	return root
}
