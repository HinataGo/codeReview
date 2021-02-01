# container 包
- list 双向链表
- ring 循环链表
- 
### 1. 可以把自己生成的Element(元素)类型值传给链表吗？
```
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveAfter(e, mark *Element)

func (l *List) MoveToFront(e *Element)
func (l *List) MoveToBack(e *Element)

# MoveBefore方法和MoveAfter方法，它们分别用于把给定的元素移动到另一个元素的前面和后面
# [comment]: <> (MoveToFront方法和MoveToBack方法，分别用于把给定的元素移动到链表的最前端和最后端)
```
- as: 不会接受，这些方法将不会对链表做出任何改动。
因为我们自己生成的Element值并不在链表中，
所以也就谈不上“在链表中移动元素”。
更何况链表不允许我们把自己生成的Element值插入其中

- reason: 
  - List包含的方法中，用于插入新元素的那些方法都只
  接受interface{}类型的值。这些方法在内部会使用Element值，
  包装接收到的新元素
  - 对于list 链起来的ele元素, 为了表明当前ele元素属于哪个list,
    在每个ele中都会存储当前属于的list指针; 这样就避免了内存泄露风险
    (为了避免直接使用我们自己生成的元素，主要原因是避免链表的内部关联，遭到外界破坏)
    
```bazaar
func (l *List) Front() *Element
func (l *List) Back() *Element

func (l *List) InsertBefore(v interface{}, mark *Element) *Element
func (l *List) InsertAfter(v interface{}, mark *Element) *Element

func (l *List) PushFront(v interface{}) *Element
func (l *List) PushBack(v interface{}) *Element

# PushFront、PushBack、InsertAfter、InsertBefore这些方法接收的参数都是interface{}，
# 最后调用内部方法的时候，都会包装成&Element{Value: v} 这样去调用

# Front和Back方法分别用于获取链表中最前端和最后端的元素
# InsertBefore和InsertAfter方法分别用于在指定的元素之前和之后插入新元素
# PushFront和PushBack方法则分别用于在链表的最前端和最后端插入新元素
# 这些方法都会把一个Element值的指针作为结果返回，
# 它们就是链表留给我们的安全“接口”。拿到这些内部元素的指针，我们就可以去调用前面提到的用于移动元素的方法了

```

#### 1.1 为什么链表可以做到开箱即用？
- 0. root Element // sentinel list element, only &root, root.prev, and root.next are used(链表 有哨兵)
- 1. List和Element都是结构体类型
- 2. 结构体类型有一个特点，那就是它们的零值都会是拥有特定结构，
     但是没有任何定制化内容的值，相当于一个空壳。值中的字段也都
     会被分别赋予各自类型的零值。
    - 零值: 只做了声明，未做初始化的变量 被给予 的 缺省值
    - var a [2]int  声明的变量a的值, 将会是一个包含了两个0的整数数组。
    - var s []int   声明的变量s的值将会是一个[]int类型的、值为nil的切片
- 3. var l list.List声明的变量l的值, 这个零值将会是一个长度为0的链表。
    - (长度为 0，但会有哨兵节点。)
- 4. 这个链表持有的根元素也将会是一个空壳，其中只会包含缺省的内容。那这样的链表我们可以直接拿来使用吗？
    - yes, 这就是开箱即用.
- 延迟初始化 机制:(这是开箱即用原因) | 初始化操作延后，仅在实际需要的时候才进行
    - 1. 它可以分散初始化操作带来的计算量和存储空间消耗
         (如果我们需要集中声明非常多的大容量切片的话，
         那么那时的 CPU 和内存空间的使用量肯定都会一个激增，
         并且只有设法让其中的切片及其底层数组被回收，内存使用量才会有所降低。)
    - 2. 如果数组是可以被延迟初始化的，那么计算量和存储空间的压力就可以被分散到实际使用它们的时候。
         这些数组被实际使用的时间越分散，延迟初始化带来的优势就会越明显。
    - 3. 同理 Go 语言的切片就起到了延迟初始化其底层数组的作用
    - 4. 链表实现中，一些方法是无需对是否初始化做判断的(Front方法和Back方法，一旦发现链表的长度为0, 直接返回nil)
    - 5. 用于删除元素、移动元素，以及一些用于插入元素的方法中，
         只要判断一下传入的元素中指向所属链表的指针，是否与当前链表的指针相等就可以了
        - 如果不相等，就一定说明传入的元素不是这个链表中的，后续的操作就不用做了。 
          反之，就一定说明这个链表已经被初始化了
    - 6. 链表的PushFront方法、PushBack方法、PushBackList方法以及PushFrontList方法总会先判断链表的状态，
         并在必要时进行初始化，这就是延迟初始化(只有在空list 中插入时才会初始化)我们在向一个空的链表中添加新元素的时候，
         肯定会调用这四个方法中的一个，这时新元素中指向所属链表的指针，
         一定会被设定为当前链表的指针。所以，指针相等是链表已经初始化的充分必要条件。
#### 1.2 Ring & List 区别?
- ring 是循环列表或环的元素。
  - 戒指没有起点或终点；指向任何环元素的指针都用作整个环的引用。
  - 空环表示为nil环指针。环的零值是一个零元素的环。
- 区别
  1. 初始化值不同:
    - 1. var r ring.Ring语句声明的r将会是一个长度为1的循环链表，
         而List类型的零值则是一个长度为0的链表(哨兵)List中的根元素不会持有实际元素值，因此计算长度时不会包含它
    - 2. 创建并初始化一个Ring值的时候，我们可以指定它包含的元素的数量，但是对于一个List值来说却不能这样做（也没有必要这样做）。循环链表一旦被创建，其长度是不可变的。这是两个代码包中的New函数在功能上的不同




