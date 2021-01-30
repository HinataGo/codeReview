# error
- package error
  - /errors
    - 1. 用指针设计对象的判断等值与否
            ```bazaar
            // 底层源码 
            // 这里注意 return 是一个 error的地址, 不要试图判断两者的error内容是否相同
            // 防止在项目中与他人代码 产生冲突的error定义 ,这里使用 & 取地址,以后写error也是
            func New(text string) error {
            return &errorString{text}
            }
            
            // errorString is a trivial implementation of error.
            type errorString struct {
            s string
            }
            
            func (e *errorString) Error() string {
            return e.s
            }
            
            ```
         // 正确使用示例
         看代码示例
    - 2. err 必须先判断 != nil 才能用, err 必须空的才是正常的,否则就错
            ```bazaar
               v, err := handler()
               if err != nil{
                    return
               } 
               // TODO v
            ```
    - 3. panic != java 异常,不一样
        - panic == fatal error 挂了,代码无法运行
        - 不要随便写
        - Request-Driven 请求驱动
        - 打印出异常,手动处理,不要包panic 错误
        - 一般也不会挂,挂了说明代码有问题
        ```bazaar
         // 野生gorountine 
         // 除了问题 ,recover不住的    
           go func(){
           }
      
         // 下面是不推荐示例.直接理解为不好的方式
         func GoXxx(x func()){
            if err := recover(); err != nil{
            }
            go x()
         }
         // 加sync 包解决,讲 goroutine 放在channl 里面,建立一个池 ,处理go 程
         // 不要来一个创建一个go 程
         package sync
         type Message struct{
         }
         ch <- &Message{}
        ```
      - 4. 强依赖 & 弱依赖
        ```bazaar
        
        ```
      - 5. 返回值不使用 指针 直接带error就好,查不到数据也认为是异常
        ```bazaar
        
        ```