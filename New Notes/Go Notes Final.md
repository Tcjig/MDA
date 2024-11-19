# Go Notes



## 基本变量与类型

### 输出格式化

| 格式化词 | 含义           |
| :------- | -------------- |
| %d       | 整数           |
| %f       | 浮点数         |
| %s       | 字符串         |
| %t       | 布尔值         |
| %v       | 适用于所有类型 |

### 变量定义：

+ var定义：
  + ``var 变量 数据类型 = 赋值``
  + ``var 变量 = 赋值``

+ 简短定义：``变量 := 赋值``

+ 批量定义：
  + ``var 变量1;变量2 = 赋值1,赋值2``
  + ``变量1,变量2 := 赋值1;赋值2``
+ 常量定义：``const 变量 = 赋值``

+ 全局变量定义：

  ```
  var (
  	变量1 = 赋值
  	变量2 = 赋值
  )
  ```

### 基本数据类型

查看变量类型：``%T``

查看字节：``unsafe.Sizeof(变量名)``

整数类型：

| 类型   | 有无符号 | 占用存储空间 | 表数范围     |
| :----- | -------- | ------------ | ------------ |
| int8   | 有       | 1            | $-2^{7}$~$2^{7}-1$ |
| int16  | 有       | 2            | $-2^{15}$~$2^{15}-1$ |
| int32/rune | 有       | 4            | $-2^{31}$~$2^{31}-1$ |
| int64  | 有       | 6            | $-2^{63}$~$2^{63}-1$ |
| uint8/byte | 无       | 1            | 0~255 |
| uint16 | 无       | 2            | 0~$2^{16}-1$ |
| uint32 | 无       | 4            | 0~$2^{31}-1$ |
| uint64 | 无       | 8            | 0~$2^{63}-1$ |

浮点类型：

| 类型            | 存储空间 | 表数范围             |
| --------------- | :------- | -------------------- |
| float32         | 4        | -3.403E38~3.403E38   |
| float64(精度高) | 8        | -1.798E308~1.798E308 |

字符类型：

+ 若想打印你想输出的字符，必须采用格式化输出：``fmt.Printf("%c",变量)``

+ 若不采用格式化输出，打印出来会显示该字符对应的ASCII值

转义字符：

| 转义符 | 含义                                     |
| ------ | ---------------------------------------- |
| \b     | 往回退一格内容                           |
| \\n    | 换行                                     |
| \r     | 光标回到本行开头，后续输入会替换原有字符 |
| \t     | 8个字符为一个制表符                      |
| \\"    | 双引号                                   |
| \\'    | 单引号                                   |
| \\\    | 反斜杠                                   |

布尔类型**(bool)**：

+ 只允许取**true**和**false**
+ 占1个字节

字符串类型(string)：

+ 字符串中字符的值不可改变
+ 若字符串中有特殊字符，可使用``` ` ` ```
+ 字符串可用``+``连接
+ 字符串过长时，可以按回车换行，但要保证**加号留在每行末尾**

### 基本数据类型转为string

显示转换：``转换类型(变量)``，括号里的变量类型不变

基本类型转为string类型：

+ ``fmt.Sprintf("参数",变量)``
+ 使用strconv包的函数**(用之前导入包)**：
  + 整形变量转：``strconv.FormatInt(int64(变量),10)``,这里的10代表10进制
  + 浮点型变量转：``strconv.FormatFloat(变量,'f',保留小数点几位,变量类型)``
  + 布尔型变量转：``strconv.FormatBool(变量)``

### string类型转为基本数据类型

使用strconv包**(ParseBool返回值有两个：value bool,err error)**：

+ 转成布尔型：``变量, _ = strconv.ParseBool(字符串变量)``

+ 转成整数型：``变量, _ = strconv.ParseInt(字符串变量,10,64)``
+ 转成浮点型：``变量, _ = strconv.ParseFloat(字符串变量,64)``



## 复杂数据类型

### 指针(内存地址)

定义指针变量：``var 指针变量 *指针类型 = &变量``

使用指针：``*指针变量``

符号理解：

+ &：取内存地址
+ *：根据地址取值

<span style="color:red">细节：</span>

+ 可通过指针改变指向值：``*指针变量 = 赋值``
+ 指针变量接收的一定是地址值
+ 指针变量的地址不可以不匹配：比如浮点型指针接受整形变量会报错

### 标识符

下划线_：为空标识符，作为占位符使用

起名规则：

+ 不要和标准库起冲突
+ 若变量名、函数名、常量首字母大写，则可以被本包和其它包访问；若首字母小写，则只能在本包使用



## 运算符

### 算数运算符

+ ++、--操作只能写在变量后面
+ ++、--操作不能参与到运算中

### 获取用户终端输入

Scanln输入：``fmt.Scanln(&变量)``，必须传入地址，否则不会修改

Scanf输入：``fmt.Scanf("格式化符号i",&变量)``，这个用于用户输入多个值



## 流程控制

### 分支结构

#### if

```
if 条件表达式 {
	语句
} else if 条件表达式 {
	语句
} else {
	语句
}
```

#### switch

```
switch 表达式 {
	case 值 :
		语句
	case 值 :
		语句
	default:
		语句
}
```

<span style="color:red">细节：</span>

+ case后的各个值的数据类型，必须和switch的表达式数据类型一样

+ switch后面放常量值、变量、一个有返回值的函数都可以

+ case后面可以有多个值，用 **,** 隔开
+ case后面不用带break
+ switch穿透：在case语句后增加**fallthrough**，将会继续执行下一个case

### 循环结构

#### for

```
// for初始表达式不能用var定义变量，用:=
for 初始表达式 ; 布尔表达式 ; 迭代因子 {
	循环体
}
```

#### for range(遍历数组、切片、字符串、map、通道)

```
// 字符串索引值被变量1接收，具体数值被变量2接收
for 变量1 , 变量2 := range 字符串变量 {
	语句
} 
```

for range和for在遍历字符串上的区别：for range可以遍历汉字，for不可以

### 关键字

#### break

作用：结束**离break最近**的循环

#### 标签

```
// 如果if符合条件了，它将结束被加了标签的外侧循环，而非内侧循环
// 若设置了标签未使用，将会报错
lable:
for ... {
	for ... {
		if ... {
			break lable
		}
	}
}
```

#### continue

作用：结束本次循环，继续下一次循环

#### return

作用：结束当前函数



## 函数

<span style="color:red">值传递：</span>基本数据类型和数组默认都是值传递，在函数内修改，不会影响原来的值

<span style="color:red">引用传递：</span>如果希望在函数内的变量修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量

### 固定参数

```
// 有返回值：一个返回值
func 函数名 (变量 数据类型) (返回值类型) {
	语句
	return 返回值
}

// 有返回值：多个返回值
func 函数名 (变量 数据类型) (返回值类型,返回值类型) {
	语句
	return 返回值1,返回值2
}
//无返回值
func 函数名 (变量 数据类型) (返回值类型) {
	语句
}
```

有无返回值区别：有返回值，则返回一个具体值；无返回值，则是输出函数内的语句

多个返回值：则在main函数中需要用**多个变量接收多个返回值**，如果有**某个返回值不需要**，则用**下划线_**替代,例如``变量,_ := 函数名(赋值,赋值)``

### 可变参数

```
// 定义含有任意数量的参数
func 函数名 (args...数据类型) {
	// 函数内部处理可变参数的时候，将可变参数当做切片处理
}
```

### 对函数返回值命名(不用对应返回值顺序)

```
// 这样它会自己对应返回值，无需按顺序在return后写返回值
func 函数名 (形参1 数据类型,形参2 数据类型) (变量1 返回值类型,变量2 返回值类型) {
	变量1 := 值1
	变量2 := 值2
	return
}
```

### 小知识点

函数也是一个类型可以赋给变量：``变量 := 函数名``，则此时通过调用变量等同于调用函数

函数可以作为一个形参：``func 函数名 (变量 数据类型,函数名 函数数据类型)``

自定义数据类型(等同于给已知数据类型起别名)：``type 自定义名称 数据类型``

### 内存分析

栈：存放基本数据类型**(先入后出)**

堆：存放复杂数据类型

代码区：存放代码本身

函数运行解析：每一个函数运行时会在栈里形成一个**本函数的栈帧**，而变量会在栈帧里面开辟一个空间，当函数执行完毕时，**栈帧被释放掉**

### 包(将不同函数放在不同源文件中)

作用：解决同名问题：若想定义一个同名函数，在一个文件中是不可以的，此时可以用不同包来区分

调用其它包函数：``包名.函数名()``

<span style="color:red">细节：</span>

+ 包的声明和包所在文件夹要同名，也可不同名
+ 一个目录内不能有重复函数
+ 一个目录下的**同级文件包的声明要一致**，否则会出错

### init函数

作用：初始化函数，进行一些初始化操作；每一个源文件可以包含一个init函数，该函数会在main函数执行前被调用

```
func init() {
	语句
}
```

### 匿名函数

作用：只调用一次的函数，没有必要单独定义一个函数

```
// 在定义匿名函数时直接就调用了
func (变量 数据类型) (返回值类型) {
	return 返回值
}(赋值) // 相当于直接调用,里面也可以不赋值
```

### 闭包

定义：返回的匿名函数+匿名函数以外的变量组成闭包

本质：闭包本质是一个匿名函数，只是这个函数引入了外部的变量或参数

```
// 实例：累加函数
// sum为匿名函数之外的变量
func getSum() func (int) int {
	var sum int = 0 
	return func (num int) int { // 这里为返回的匿名反函数
		sum = sum + num
		return sum
	}
}

func main() {
	f := getSum()
	fmt.Print(f(1)) // 输出1
	fmt.Print(f(2)) // 输出3
	fmt.Print(f(3)) // 输出6
	// 上面的输出知道匿名函数中引用的外部变量会一直保存在内存中，可以一直使用
}
```

<span style="color:red">注意：</span>

+ 闭包中使用的变量或参数会一直保存在内存中，因此会一直使用**(对内存消耗大)**
+ 不使用闭包的缺点：想保留的值**不可以反复使用**
+ 闭包应用场景：闭包可以保留上次引用的值，传入一次就**可以反复使用**了

### defer关键字

作用：在函数执行完毕后，及时**释放资源**

```
// 实例
func main() {
	fmt.Println(add(30,60))
}

func add(num1 int,num2 int) (int) {
	//程序会将defer后的语句压入一个栈中，然后执行函数后面的语句
	//在压入栈中时，会将相关值同时拷贝进栈中，不会随函数后面变化而变化
	//在函数执行完后，从栈中取出语句执行
	defer fmt.Print("num1=",num1)
	defer fmt.Print("num2=",num2)
	var sum int = num1 + num2
	fmt.Println("sum=",sum)
	return sum
}
// 执行结果
sum=90
num2=30
num1=60
90
```

应用场景：用于你想关闭某个使用的资源，defer有延迟执行机制，所以当你用完defer后他会随时关闭

### 系统函数

#### 字符串函数(使用strconv和strings时要导入包)

+ 统计字符串长度**(按字节统计)**：``len(字符串变量)``

+ 字符串遍历**(for range也是一种方法)**：``r := []rune(字符串变量)``，这里将字符串转成了**切片**，当转完后，**可以正常用for进行遍历**

+ 字符串转整数：``变量,_ := strconv.Atoi("字符串")``，因为这里它有可能返回两个值，一个是转换完成的值，另一个是报错值，所以要用**err或下划线_**接收第二个值

+ 整数转字符串：``字符串变量 = strconv.Itoa(整数)``

+ 查找子串是否在指定的字符串中：``strings.Contains("母串","子串")``

+ 统计一个字符串有几个指定字符串：``变量 = strings.Count("母串","子串")``，他这里返回的是一个整数，有几个指定子串出现在母串，整数就为多少

+ 不区分大小写的字符串比较：``bool = strings.EqualFold("字符串1","字符串2")``，返回true或false

+ 返回子串在字符串第一次出现的索引值，如果没有返回-1：``strings.Index("母串","子串")``

+ 字符串替换：

#### 日期和时间相关函数

Now函数：

+ 具体时间(包含年月日时分秒等)：``time.Now()``
+ 年：``now.year()``
+ 月：``now.Month()``
+ 日：``now.Day()``
+ 时：``now.Hour()``
+ 分：``now.Minute()``
+ 秒：``now.Second()``

格式化输出年月日时分秒：

```
// 实例
// 这个方法可以得到字符串，以便后续使用
datestr := fmt.Sprintf("年月日：%d-%d-%d 时分秒：%d-%d-%d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.m,Second())
fmt.Println(datestr)
// 执行结果
年月日：... 时分秒：...
```

按照指定格式输出：``变量 := now.Format("自定义数字")``

<span style="color:red">注意：</span>需要导入time包

#### 内置函数(不需要导包，可直接用)

统计字符串、数组、切片、通道等长度：``len(变量)``

统计容量：``cap(变量)``

分配内存：``new(数据类型)``



## 错误处理

作用：当你代码中某一块出错时，它会导致整个代码无法运行，而运用错误处理机制，它可以帮你把**错误的信息变成一个返回值**，从而可以将错误的地方进行输出，以便整体代码的正常运行以及可以**看到错误反馈**

defer+recover捕获错误：

```
// 实例
func main() {
	test()
	fmtPrint("正常执行")
}

func test() {
	defer func() {
		err := recover() // 使用err变量接受错误
		if err != nil {
			fmt.Println("错误是：",err) // 这里会将错误信息打印出来
		}
	}() // 这里的意思是匿名函数被直接调用了
	num1 := 10
	num2 := 0
	result := num1 /num2
	fmt.Println(result) // 这里可以看到10/0是错误的，因此错误被err变量接收了
}

// 执行结果
错误是：...
正常执行
```

自定义错误：



## 数组(存储相同类型的数据)

### 一维数组

定义数组：``var 数组名 [数组长度]数据类型``

数组优点：访问、查询、读取速度快

数组遍历：

```
for 变量1,变量2 := range 数组名 {
	语句
}
```

<span style="color:red">注意：</span>变量1接受的是**数组的索引**，变量2接收的是**数组的值**，若有的值**不想接收**，可以用**下划线_**接收

数组初始化(举例)：

+ ``var arr [3]int = [3]{1,2,3}``
+ ``var arr = [3]int{1,2,3}``
+ 不确定长度：``var arr = [...]int{1,2,3}``
+ 不确定长度且给固定索引赋值：``var arr = [...]int{0:1,1:2,2:3}``

数组注意事项：

+ 数组默认为值传递、值拷贝

+ 若想在其他函数中修改外部数组，可以用指针(引用传递)

  ```
  // 实例
  func main() {
  	var arr = [3]int[1,2,3]
  	test(&arr) //此时arr[0]的值被改为4
  }
  func test(arr *[3]int) {
  	(*arr)[0] = 4
  }
  ```

### 二维数组

定义数组：``var 数组名 [值1][值2]数据类型``，数值1为**大格子**，数值2为在大格子里的**小格子**

数组初始化：``var 数组名 [值1][值2]数据类型 = [值1][值2]数据类型{{...},{...}...}``

数组遍历：

```
// 实例
var arr [2][2]int = [2][2]int{{1,2},{3,4}}
for a,b := range arr {
	for c,d := range b {
		... // 填写输出语句即可
	}
}
```



## 切片(slice)

### 切片的定义及遍历

解释：切片是建立在数组的基础上的，切片为截取数组的某个片段，这个片段也可以是整个数组

切片包含的数据结构：指向底层数组的指针、切片的长度、切片的容量

定义切片：

+ ```
  // 因为切片是以数组为基础的，所以要先定义一个数组
  var arr [6]int = [6]int{1,2,3,4,5,6}
  // 这里定义的切片截取数组中索引为1到索引为3(不包含索引为3)的内容
  var slice []int = arr[1:3]
  fmt.Println(slice)
  
  //执行结果
  [2 3] // 因为不包含末尾索引，所以值4不输出
  ```

+ make内置函数定义切片：``var 切片名 := make([]数据类型,切片长度,切片容量)``

+ 直接定义切片**(类似于make内置函数)**：``切片名 := []数据类型{值1,值2,...}``

切片遍历：

```
// 实例
slice := make([]int,4,20)
for a,b := range slice {
	... // 填写输出语句即可
		// 只不过这里输出的slice的值都为0
		// 因为没有给他赋值,而int类型默认值为0
}
```

切片的注意事项：

+ 定义切片后不可以直接使用(除非使用make内置函数定义)，需要让其引用到一个数组
+ 切片中截取的数组片段，初始索引值**可取到**，末尾索引值**不可取到**
+  切片可以继续切片

<span style="color:red">切片相较于数组的优点：</span>

+ 切片的容量为动态变化，不是固定的
+ 切片是对数组一个连续片段的引用，因此切片是一个引用类型，这也就意味着你**通过修改切片的值，初始数组和再次切片的值也会被改变**

### 切片的动态增长(append)

通过append给切片增加数组和切片：

```
// 实例
var arr [6]int = [6]int{1,2,3,4,5,6}
var slice1 []int = arr[1:4]
slice1 = append(slice1,100,200) // 第一种追加数组方式：使用原先的切片接收
slice2 = append(slice1,100,200) // 第二种追加数组方式：使用新切片接收
fmt.Println(slice1)
fmt.Println(slice2)

slice3 := []int{300,400}
slice1 = append(slice,slice2...) // 追加切片，后面的三个点必须写，表示追加的是一个切片
fmt.Println(slice)
// 执行结果
[2 3 4 100 200]
[2 3 4 100 200]
[2 3 4 100 200 300 400] 
```

原理：如果只是单纯用append追加数组，后面不用新切片或原先的切片接收，是无法追加的，这是因为append会在底层新创建一个数组，而原先的切片和append追加的数组都被拷贝到新数组里了

<span style="color:red">注意：</span>

+ 第一种追加数组方式是常用的
+ 追加切片时，后面的三个点必须写，表示追加的是一个切片

### 切片的拷贝(copy)

```
// 实例
var a []int = []int{1,2,3,4,5,6}
var b []int = make([]int,10)
copy(b,a) // 将a中对应数组中元素内容复制到b中对应数组中去
fmt.Println(b)

//执行结果
[1,2,3,4,5,6,0,0,0,0]
```

<span style="color:red">注意：</span>这里修改b或a切片对另外的一个切片都是没有影响的，因为它们底层都各自对应一个数组



## 映射(map)

### map的定义及使用

解释：它将**键值对**相关联，可以通过**键(key)**来**获取**对应的**值(value)**

定义map：``var 变量名 map[数据类型]数据类型``，这里只是单纯的声明了map，因此map内存没有分配空间，只有**通过make函数对它初始化**，才会分配空间

使用map**(make函数)**：

```
var a map[int]string // []里对应的是key，[]后面对应的是value
a = make(map[int]string,10) // 这里的10显示map可以存入10个键值对
// 将键值对存入map中
a[1] = "first"
a[2] = "second"
a[3] = "third"
fmt.Println(a)

// 执行结果：它将无序输出这些键值对
map[2:second 1:first 3:third]
```

<span style="color:red">注意：</span>

+ map集合在使用前一定要make
+ map的键值对是无序的
+ key是不可以重复的，若遇到重复，后一个value会替换前一个value
+ value可以重复
+ make函数的第二个参数可以省略，也就是上述代码10的位置可以不写，但它默认只分配一个内存

### map创建方式

```
// 方式一
var a map[int]string
a = make(map[int]string,10)

// 方式二
b := make(map[int]string)

// 方式三
c := map[int]string {
	1 : "first",
	2 : "second",
}
```

### map的操作

#### 增加和更新

```
a := make(map[int]string)
// 增加
a[1] = "first"
a[2] = "second"
a[3] = "third"

// 修改
a[2] = "two"

// 删除
delete(a,1)

// 查找
c,d := a[2] // 用c接收找到的值，d接收布尔类型的值
// d接收的值：找到返回true，找不到返回false
```

#### 单层map遍历

```
// 实例
a := make(map[int]string)
b[1] = "first"
b[2] = "second"
b[3] = "third"

for b,c := range a {
	fmt.Printf("key：%v value：%v\n",b,c)
}

// 执行结果
key：1 value：first
key：2 value：second
key：3 value：third
```

#### 叠加map遍历

```
// 实例
// key对应string类型，value对应map[int]string
// 而value对应的map里面的key对应int类型，里面的value对应string类型
a := make(map[string]map[int]string)

// 赋值
a["数字"] = make(map[int]string,3)
a["数字"][1] = "first"
a["数字"][2] = "second"
a["数字"][3] = "third"

for k1,v1 := range a {
	fmt.Println(k1)
	for k2,v2 := range v1 {
		fmt.Printf("数字：%v 字母：%v\n",k2,v2)
	}
}

// 执行结果
数字
数字：1 字母：first
数字：2 字母：second
数字：3 字母：third
```



## 文件的操作(os)

### 文件的基本操作

打开文件：``变量,err := os.Open("文件路径");``，这里变量接收返回的**文件指针**，err接收是否**打开成功**的信息

关闭文件：``err := file.Close();``，err接收是否**关闭成功**的信息

<span style="color:red">注意：</span>要导入os包

### IO流

作用：程序和数据源之间的桥梁，对**文件内部数据**进行操作

#### IO读取文件(程序从源文件接收信息)

一次性读取文件，适用于文件较小：

```
// 这里不用写文件的开关操作
变量,err :=ioutil.ReadFile("文件路径")
```

<span style="color:red">注意：</span>如果输出的文件信息是一堆码，使用string转换即可

带缓冲读取文件，适用于文件较大：

```
// 实例
// 导包
import (
	"fmt"
	"os" // 负责打开关闭文件
	"bufio" // 创建程序和文件之间的通道
	"io" // 读取文件
)

// 这里要写文件开关操作
// 打开文件
变量1,err := os.Open("文件路径")
// 当函数退出时，让文件关闭，防止内存泄漏
defer 变量1.Close()
// 创建一个流
reader := bufio.NewReader(变量1)
// 读取操作
for {
	变量2,err := reader.ReadString('\n') // 读取到一个换行就结束
	if err == io.EOF { // io.EOF表示已经读到文件的结尾
		break
	}
	// 如果没有读到文件结尾的话，就正常输出内容即可
	fmt.Println(变量2)
}
```

<span style="color:red">注意：</span>使用os、bufio、io等包时一定要**导入包**

#### IO写入(程序往源文件写入信息)

```
// 实例
// 导包
...
// 打开文件
变量1,err := os.OpenFile("文件路径",操作文件模式,文件权限) // 如果有多个模式，用|进行连接,后面那个文件权限在windows系统下可以随便写，不生效
if err != nil { // 文件打开失败，则程序结束
	fmt.Println("文件打开失败")
	return
}
// 文件关闭
defer 变量1.Close()
// 写入文件操作：先创建IO流，在进行缓冲输出流
变量2 := bufio.NewWriter(变量1)
// 写单个内容
变量2.WriteStringt("内容")
// 写多个内容
for a := 0;a < 10;i++ {
	writer.WriteString("内容")
}
// 流带缓冲区，刷新数据，将IO管内数据写入文件中
变量2.Flush()
```

#### IO复制

```
// 定义源文件
变量1 := "文件路径"
// 定义目标文件(被复制到的文件)
变量2 := "文件路径"
// 对文件进行读取
变量3,err := ioutil.ReadFile(变量1)
if err != nil { // 文件打开失败，则程序结束
	fmt.Println("读取有问题")
	return
}
// 写入文件
err = ioutile.WriterFile(变量2,变量3,文件权限)
if err != nil {
	fmt.Println("写出失败")
}
```



## 协程和管道

### 基本概念(程序、进程、线程、协程)

程序：为了完成特定任务，用某种语言编写的一组指令的集合

进程：是程序的一次执行过程，指正在运行的一个程序,进程是动态的,有生命周期

线程：是进程的进一步细化，是一个程序内部的一条执行路径，若一个进程同一时间执行多个线程，就是支持多线程的

协程：它可以在执行A函数时，随时中断，去执行B函数，然后继续中断，回去继续执行A，这一切换**并不是函数调用**，过程类似多线程，但协程中只有一个线程在执行**(本质是单线程)**

### 协程(goroutine)

#### 启动单协程

```
// 实例
import(
	"fmt"
	"time"
)
func test() {
	for i := 1;i <= 10;i++ {
		fmt.Println("hello b")
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}
func main() { // 主线程
	go test() // 开启一个协程
	
	for i := 1;i <= 10;i++ {
		fmt.Println("hello a")
		time.Sleep(time.Second)
	}
}
```

主线程和协程执行流程**(主死随从)**：当主线程先结束了，即使协程没有执行完也会结束

#### 启动多个协程：

```
// 实例：使用匿名函数
// 导包
...
func main() {
	for i := 1;i <= 5;i++ {
		go func(n int) {
			fmt.Print;n(n)
		}(i) // 如果不穿i它可能打印出来都是一个数，因为这里匿名函数+外部变量组成了闭包
	}
	time.Sleep(time.Second * 2) // 阻塞两秒
}

// 执行结果：它将无序打出1，2，3，4，5
```

#### 使用WaitGroup控制协程退出

作用：WaitGroup用于等待一组线程的结束，解决主死随从的问题

```
// 实例
// 导包
import(
	"fmt"
	"sync" // 用于使用WaitGroup
)
var wg sync.WaitGroup // 只定义无需赋值
func main() {
	// 启动五个线程
	for i := 1;i <= 5;i++ {
		wg.Add(1) // 协程开始时进行加一操作
		go func(n int) {
			defer wg.Done() // 协程完成时进行减一操作
			fmt.Println(n)
		}(i)
	}
	// 主线程一直在阻塞，何时wg减为0就停止
	wg.Wait()
}
```

#### 锁(多个协程操纵同一个数据)

作用：确保一个协程在执行逻辑时另外的协程不执行

##### 互斥锁(适用于读写次数没有明显区别)

加锁：``lock.Lock()``

解锁：``lock.Unlock()``

```
// 实例：定义两函数实现加一，减一操作，得出最后的结果为0
// 导包
import(
	"fmt"
	"sync"
)
var totalNum int
var wg sync.WaitGroup
var lock sync.Mutex

// 实现加一操作
func add() {
	defer wg.Done()
	for i := 0;i < 1000;i++ {
		lock.Lock()
		totalNum += 1
		lock.Unlock()
	}
}

// 实现减一操作
func sub() {
	defer wg.Done()
	for i := 0;i < 1000;i++ {
		lock.Lock()
		
		totalNum -= 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2) // 这里有几个协程就写多少
	go add()
	go sub()
	wg.Wait()
	fmt.Println(totalNum)
i
```

##### 读写锁(适用于读次数远远多于写次数)

加锁：``lock.Rlock``

解锁：``lock.RUnlock``

```
// 实例：进行读取操作
// 导包
import(
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var lock sync.RWMutex

// 实现加一操作
func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("读取中...")
	time.Sleep(time.Second)
	fmt.Println("读取完成")
	lock.RUnlock
}
func main() {
	wg.Add(2) // 这里有几个协程就写多少
	for i := 3;i < 3;i++ {
		go read()
	}
	wg.Wait()
i

// 执行结果
读取中...
读取中...
读取中...
读取完成
读取完成
读取完成
```

<span style="color:red">注意：</span>如果只是读数据，那么这个锁不产生影响，但是当读写同时发生时，就会哦有影响

### 管道(channel)

#### 管道的基本使用

特点：

+ 管道的本质是一个队列，特性是**先进先出**
+ 自身线程安全，多协程访问时，**不需要加锁**，因此当多个协程操作同一个管道时，**不会发生资源争抢问题**
+ 管道是有类型的，一个string的管道**只能**存放string类型数据
+ 管道是**引用类型**，**必须初始化**才能写入数据，即**make后才能使用**

定义管道：``var 变量名 chan 管道数据类型``

```
// 实例
// 定义一个int类型的管道
var intChan chan int
// 通过make给管道初始化，管道可以存储三个int类型的数据
intChan = make(chan int,3)
// 向管道存放数据
intChan <-10 // 第一种存放

a := 20 // 第二种存放
intChan <-num

//在管道中读取数据
num1 := <-intChan
num2 := <-intChan

fmt.Println(num1,num2)

// 运行结果
10 20
```

<span style="color:red">注意：</span>

+ 不能存放大于容量的数据
+ 在没有使用协程的情况下，如果管道的数据已经全部取出，那么再取就会报错

#### 管道的关闭

作用：使用内置函数close可以关闭管道，关闭后，就不能再向管道写数据了，但是仍然可以从管道读取数据

定义：``close(管道名)``

#### 管道的遍历(for range)

```
// 实例
var intChan chan int
intChan = make(chan int, 10)
// 通过for循环给管道赋值
for i := 0; i < 10; i++ {
		intChan <- i
}
	close(intChan) // 在遍历前，若没有进行管道关闭，就会出现deadlock的错误
for v := range intChan {
		fmt.Println(v)
}
// 执行结果
0123456789
```

<span style="color:red">注意：</span>

+ 在遍历时，若管道**没有关闭**，则会出现**deadlock错误**
+ 在遍历时，若管道**已经关闭**，则会**正常遍历**数据，遍历完后，就会退出遍历
+ 遍历管道只能用**for range**，但赋值可以用for循环

#### 协程和管道协同工作

```
// 实例：写入和读取
// 导入包
import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup
// 写入
func readData(intChan chan int) { // 设置一个chan int类型的intChan形参
	for i := 1;i <= 3;i++ {
		intChan <-i
		fmt,Print(i)
	}
	close(intChan) // 管道关闭
}

// 读取
func readData(intChan chan int) {
	for v := range intChan {
		fmt.Print(v)
	}
}

func main() {
	intChan := make(chan int,3)
	wg.Add(2)
	go writeData()
	go readData()
	wg.Wait()
}
1 1 2 2 3 3
```

#### 声明只读、只写通道

解释：默认情况下，管道是双向的，当设置完**只读或只写**时，就**无法写入或读取**了，

只读：``var 通道名 chan <-数据类型``

只写：``var 通道名 <-chan 数据类型``

#### 管道的阻塞

解释：当只写入数据，**没有读取**时会出现阻塞，也就是deadlock

解决：有读取功能，即使**写快读慢**也没事

#### select功能

作用：解决多个管道的选择问题，也就是从多个管道中随机**公平**选择一个来执行

```
// 实例
int main() {
	intChan := make(chan int,1)
	go func() {
		time.Sleep(time.Second * 10)
		intChan <-10
	}()
	stringChan := make(chan string,1)
	go func() {
		time.Sleep(time.Second * 5)
		stringChan <-"number"
	}()
	
	select{
		case v := <-intChan :
			fmt.Println(v)
		case v := <-stringChan :
			fmt.Println(v)
		default:
			fmt.Println("防止select被阻塞")
	}
}

// 执行结果

// 不写default
number
// 写default
防止select被阻塞
```

<span style="color:red">注意：</span>

+ 这里的公平指的就是等待时间，哪个时间先执行哪个
+ case后面必须进行的是io操作，而非等值
+ default是为了防止select被阻塞，当有阻塞时，执行default的语句

#### defer+recover机制处理错误

作用：当一个协程出现错误时，防止程序崩溃，因此利用该机制捕获错误，保证主线程继续执行

```
// 实例
// 导包
func printNum() {
	for i := 1;i <= 10;i++ {
		fmt.Print(i)
	}
}
func devide() {
	defer func() { // defer+recover机制处理错误
		err := recover() {
		if err != nil{
			fmt.Println("错误被捕获",err)
			}
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2 // 这里的错误在于除数为0
	fmt.Println(result)
}
func main() {
	go printNum()
	go devide()
	time.Sleep(time.Second * 5) // 等同于wg.Add()
}

// 执行结果
错误被捕获...
12345678910
```



## 网络编程

### 基本知识

IP：每个机器都有自己的IP，根据IP进行定位

端口号(PORT)：每个应用程序都有自己唯一的端口号

通信协议：设备之间进行传输时，必须遵守的规则**(TCP/IP协议)**

TCP/IP协议：分为应用层、传输层、网络层、物理及数据链路层

关于传输层：

+ 我们采用TCP连接而不采用UDP连接，因为UDP容易丢包

+ 建立连接(三次握手)
+ 释放连接(四次挥手)

### 客户端和服务器端的连接

<span style="color:red">注意：</span>

+ 要先启动服务器端，再启动客户端

#### 创建客户端(发送信息)：Dial函数

```
// 实例：创建客户端
// 导包
import(
	"fmt"
	"net" // 所需要的网络编程都在net包里
)

func main() {
	fmt.Println("客户端启动")
	// 调用Dial函数：参数需要指定TCP协议，需要指定服务器端的IP+PORT
	conn,err := net.Dial("tcp","127.0.0.1:888") //前面的变量接收连接成功的信息，后面的变量接收连接失败的信息，前面的参数是指定的协议，后面的参数是服务器端的IP+PORT
	if err != nil {
		fmt.Println("客户端连接失败，失败原因：",err)
		return // 程序结束
	}
	fmt.Println("客户端连接成功，"conn)
}
```

#### 创建服务器端(接收信息)：Lisiten函数

```
// 实例：创建服务器端
// 导包
...
func main() {
	fmt.Println("服务器启动")
	listen,err := net.Listen("tcp","127.0.0.1:8888") // 调用listen函数：进行监听，需要指定服务器端的TCP协议和服务器端的IP+PORT
	if err != nil {
		fmt.Println("监听失败，失败原因：",err1)
		return
	}
	
	// 监听成功后，等待客户端的连接
	for{ // 循环等待客户端连接：很多客户端连接时
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("客户端等待失败，失败原因：",err)
		}else {
			// 连接成功
			fmt.Printf("等待连接成功，conn=%v，接收到的客户端信息：%v		 		     			 \n",conn,conn.RemoteAddr().String())
		}
	}
}
```

#### 发送终端数据(bufio)

```
// 实例：客户端发送数据
// 导包
import(
	"fmt"
	"net"
)
func main() {
	fmt.Println("客户端启动...")
	conn,err := net.Dial("tcp",127.0.0.1:8080)
	if err != nil {
		fmt.Println("连接失败，失败原因：",err)
	}
	fmt.Println("连接成功",conn)
	reader := bufio.NewReader(os.Stdin)
	str,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败，失败原因为：",err)
		return
	}
	fmt.Println("读取成功，正在发送...")
	n,err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("发送失败，失败原因为：",err)
		return
	}
	fmt.Println("发送成功，字节为：",n)
}
```

```
// 实例：服务器接收数据
// 导包
import(
	"fmt"
	"net"
)
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("接受失败，失败原因为：",err)
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}
func main() {
	fmt.Println("服务端启动...")
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听失败，失败原因：",err)
	}
	fmt.Println("监听成功",listen)
	defer listen.Close()
	for {
		conn,err := listen.Accept() // 是一个阻塞操作，等待新的客户端连接到服务器
		if err != nil {
			fmt.Println("连接失败，失败原因：",err)
			return
		}
	fmt.Println("连接成功",conn)
	}
	process(conn)
}
```

