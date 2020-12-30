package code

// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return check(root.Left, root.Right)
	}

}
func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return check(l.Left, r.Right) && check(l.Right, r.Left)
}

// 树 是层序遍历   --- 队列辅助   前中后序是栈辅助
func bfs(p *TreeNode) []int {
	res := make([]int, 0)
	if p == nil {
		return res
	}
	queue := []*TreeNode{p}
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res = append(res, queue[0].Val)
			queue = queue[1:]
		}
	}
	return res
}
