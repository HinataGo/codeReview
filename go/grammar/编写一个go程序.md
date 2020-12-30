# golang编写基础细节
### 命名规范: 
    - 1. func xxx_test.go (xxx.go的测试文件)
    - 2. func Ttt.go 大写开头默认public,小写开头默认private
    - 
    
### 自定义类型
- type
    ```
    type other int 
    var  one int 
    // 这里认为 other 和 one 不是同一个类型, other会被当做新的类型
    // 强类型转换
    one = int(other)
    ```
- struct
    ```
    
    ```
- os.Exit 表示返回状态
    ```
    os.Exit(-1) // 异常值
    // 带入命令行参数?
    
    ```
### 变量和常量

### 数据类型
- go支持指针类型,但是不支持指针运算
- go不支持任何隐式类型转换
- go的string初始化值为 空字符串 ,不要使用nil(零值) 去判断

### go的 ++ --
- 不区分前置后置 (只支持后置)

### if条件语句
- if var xx; condition{
    // code 
  }
- 支持多返回值
    ```
    func TestOther(t *testing.T) {
    	if _, err := aa(); err != nil {
    		t.Log("OK")
    	}else {
    		t.Log("no")
    	}
    }
    func aa() (int, *int) {
    	return 0, nil
    } 
    ```
### switch case 语句
- 支持多条件的case 
- 简化if else
    ```
     func TestSwitch(t *testing.T) {
     	for i := 0; i < 6 ; i++ {
     		switch i {
     		case 0,1:
     			t.Log("A")
     		case 2, 3:
     			t.Log("B")
     		case 4, 5, 6:
     			t.Log("C")
     		}
     
     	}
     }
    ```