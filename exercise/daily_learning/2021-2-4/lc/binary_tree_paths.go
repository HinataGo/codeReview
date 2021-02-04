package lc

import "strconv"

// 注释中写明实现逻辑和步骤 ,以及为什么需要这一步的
func binaryTreePaths(root *TreeNode) []string {
	var res []string

	// 2.. 下面是默认 root != nil ..因此这里单独判断
	if root == nil {
		return res
	}
	// 1.向res 中添加数据,,这是先将左右子树给过去
	// 自上而下
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}
	// 3. 递归逻辑,先添加左子树
	// 自下而上 从底部返回,数据,,这是需要拼装数据
	// 而后遍历数据从头开始也就是自上而下, 深度遍历的当时,,一条路走到尽头的遍历
	l := binaryTreePaths(root.Left)
	for _, v := range l {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	r := binaryTreePaths(root.Right)
	for _, v := range r {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	return res
}
