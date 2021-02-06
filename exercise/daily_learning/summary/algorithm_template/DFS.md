# 深度优先遍历 
## 递归
### 二叉
```bazaar
// 前中后
var ret []int

// 主函数
func dfs(root *TreeNode) []int {
	ret = make([]int, 0)
	pre(root)
	return ret
}
// 前序
func pre(node *TreeNode) {
	if node == nil {
		return
	}
	// 下面三个分别代表, 根 左 右,
	// 其他的遍历只需要调换这三者之间的顺序即可 (只需要看 append位置,一次是前中后序)
	ret = append(ret, node.Val)
	pre(node.Left)
	pre(node.Right)
}
```
- 换一个函数的写法
```bazaar
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func dfs(root *TreeNode) []int {
	ret := make([]int, 0)
	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)
	return ret
}
```
### N叉
```bazaar
type Node struct {
	Val      int
	Children []*Node
}

func dfs(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	// 下面两步骤顺序 目前是前序,互换为 后序遍历(没有中序)
	res = append(res, root.Val)
	// 特殊点 ,这里不要分开函数写了
	// 因为没有左右孩子区分,只有多个children ,因此这里,直接遍历左右孩子
	// 依次添加
	//
	for _, v := range root.Children {
		res = append(res, dfs(v)...)
	}
	return res
}
```