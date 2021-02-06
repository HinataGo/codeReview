# 算法模板 必背
## 常用API 操作
- 切片的取值 slice[:] --> 是[) 左开右闭 不包括扩右边界
- 参数传递，只能修改，不能新增或者删除原始数据
- leetcode 中，全局变量不要当做返回值，否则刷题检查器会报错,这里可以每次使用前初始化全局变量,这样已达到清空目的
### stack 栈
```bazaar
// 创建栈
stack:=make([]int,0)
// push压入
stack=append(stack,10)
// pop弹出
v:=stack[len(stack)-1]
stack=stack[:len(stack)-1]
// 检查栈空
len(stack)==0
```
### queue 队列
```bazaar
// 创建队列
queue:=make([]int,0)
// enqueue入队
queue=append(queue,10)
// dequeue出队
v:=queue[0]
queue=queue[1:]
// 长度0为空
len(queue)==0
```
### map 字典
- map 键需要可比较，不能为 slice、map、function
- map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由 0 变为 1
- 比较两个 map 需要遍历，其中的 kv 是否相同，因为有默认值关系，所以需要检查 val 和 ok 两个值
```bazaar
// 创建
m:=make(map[string]int)
// 设置kv
m["hello"]=1
// 删除k
delete(m,"hello")
// 遍历
for k,v:=range m{
    println(k,v)
}
```
## 标准库API
### sort
```bazaar
// int排序
sort.Ints([]int{})
// 字符串排序
sort.Strings([]string{})
// 自定义排序
sort.Slice(s,func(i,j int)bool{return s[i]<s[j]})
```
### math
```bazaar
// int32 最大最小值
math.MaxInt32 // 实际值：1<<31-1
math.MinInt32 // 实际值：-1<<31
// int64 最大最小值（int默认是int64）
math.MaxInt64
math.MinInt64
```
### copy
```bazaar
// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
// copy(待粘贴位置,复制的值) 这么记忆
copy(a[i:],a[i+1:])
a=a[:len(a)-1]

// make创建长度，则通过索引赋值
a:=make([]int,n)
a[n]=x
// make长度为0，则通过append()赋值
a:=make([]int,0)
a=append(a,x)
```
### 类型转换
```bazaar
// byte转数字
s="12345"  // s[0] 类型是byte
num := int(s[0]-'0') // 1
str := string(s[0]) // "1"
b   := byte(num+'0') // '1'
fmt.Printf("%d%s%c\n", num, str, b) // 111

// 字符串转数字
num,_ := strconv.Atoi() // 字符转数字
str   := strconv.Itoa() // 数字转字符
```