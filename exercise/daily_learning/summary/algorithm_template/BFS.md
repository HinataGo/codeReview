# 广度优先搜索 / 层序遍历
## 递归
### 二叉树
```bazaar
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
// 全局变量
var res [][]int
// 主调用函数
func levelOrder(root *TreeNode) [][]int {
// 初始化全局变量
	res = make([][]int, 0)
	bfs(root, 0)
	return res
}
// 层序遍历模板
func bfs(root *TreeNode, level int) [][]int {
	if root == nil {
		return res
	}
	if level == len(res) {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}
	// 递归调用左右孩子
	res = bfs(root.Left, level+1)
	res = bfs(root.Right, level+1)
	return res
}
```
### N叉树
```bazaar
type Node struct {
	Val      int
	Children []*Node
}
// 主函数
func levelOrder(root *Node) [][]int {
	res = make([][]int, 0)
	bfs(root, 0)
	return res
}
func bfs(root *Node, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	// 与二叉树的层序遍历不同之处在于这里
	// 二叉是 不等情况下 append root.val 给 res[level] ,N叉是无论如何都给
	// N叉是只需要遍历自己的Children然后递归调用level+1 即可,因为没有左右孩子之分
	
	res[level] = append(res[level], root.Val)
	for _, v := range root.Children {
		bfs(v, level+1)
	}
}
```
## 非递归