package lc

func hasPathSum(root *TreeNode, targetSum int) bool {
	// if root == nil{
	//     return targetSum == 0
	// }
	// bug [1,2] 1 这里root为根节点,,减去之后呢发下sum = 0 ,
	// 但是他不是叶子结点路径,仅仅是一个根节点,,因此这里只算falsefalse
	//  必须强制判断 这个结点是否为叶子结点
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
