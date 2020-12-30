package code

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

// 前序遍历
// func per(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	per(root.Left)
// 	per(root.Right)
// }
