package note

import (
	"fmt"
	"gonote/util"
	"sync"
)

var v = 100

func EscapedCharacters() {
	v = 200
	fmt.Println("1 双引号")
	fmt.Println("图灵祖师说：\"我不知道我是梦见自己是机器的图灵，还是梦见自己是图灵的机器")
	fmt.Println("\n2 反斜线")
	fmt.Println("电子邮件说\\项目已取消\\清理文档的时候\\我哭了")
	fmt.Println("\n3 回车")
	// 一般配合 \n\r 使用
	fmt.Println("我的感官很悠闲，我的精神自由地按照它自己的直觉前进\r我的程序如同是自己在写自己")
}

const (
	Version int = 100
)

// 2.2 变量与常量
func VariablesAndConstants() {
	fmt.Println("\n1 变量")

	var v1 int
	v1 = 1
	var v2 int = 2
	var v3 = 3
	v4 := 4
	fmt.Printf("v1=%v,v2=%v,v3=%v,v4=%v\n", v1, v2, v3, v4)
	var (
		v5     = 5
		v6 int = 6
		v7 int
	)
	fmt.Printf("v5=%v,v6=%v,v7=%v\n", v5, v6, v7)

	fmt.Println("\n2 常量")

	const (
		c1 = 8
		c2 = iota // 当前行数，从0开始
		c3        // 默认值为上 行的值
		c4
		c5 = 12
		c6
	)

	fmt.Printf("c1=%v,c2=%v,c3=%v,c4=%v,c5=%v,c6=%v\n", c1, c2, c3, c4, c5, c6)

}

// 2.3 基本数据类型
func BasicDataTypes() {
	fmt.Println("\n2.3.1 基本数据类型")
	var (
		n1        = 0b0101
		n2 int8   = 0o77
		n3 uint16 = 0xAF
	)
	fmt.Printf("n1=%v,type is %T\n", n1, n1)
	fmt.Printf("n2=%v,type is %T\n", n2, n2)
	fmt.Printf("n3=%v,type is %T\n", n3, n3)

	fmt.Println("\n2.3.2 浮点型数据类型")
	var (
		f1         = 1.0
		f2 float32 = 1
	)
	fmt.Printf("f1=%v,type is %T\n", f1, f1)
	fmt.Printf("f2=%v,type is %T\n", f2, f2)

	fmt.Println("\n2.3.3 数值型数据类型转换")
	n2 = int8(n3)
	fmt.Printf("n2=%v,type is %T\n", n2, n2)

	fmt.Println("\n2.3.4 字符型数据类型")

	var (
		c1 byte
		c2      = '0'
		c3 rune = '中'
	)
	fmt.Printf("c1的码值=%v,这个码值对应的字符是%c,type is %T\n", c1, c1, c1)
	fmt.Printf("c2的码值=%v,这个码值对应的字符是%c,type is %T\n", c2, c2, c2)
	fmt.Printf("c3的码值=%v,这个码值对应的字符是%c,type is %T\n", c3, c3, c3)

	fmt.Println("\n2.3.5 布尔型数据类型")
	var bool bool = true
	fmt.Printf("bool=%v,type is %T\n", bool, bool)

	fmt.Println("\n2.3.6 字符串型数据类型")
	var s1 string = "hello"
	fmt.Println(s1 + "world")
	fmt.Println(s1, "world")
	fmt.Println(len(s1))
	// 预格式化的字符串
	var s2 = `	var (
		c1 byte
		c2      = '0'
		c3 rune = '中'
	)`
	fmt.Println(s2)

}

func increase(n *int) {
	*n++
	fmt.Printf("increase 结束时, n=%v,内存地址为%v,指向的原始值为%v\n", n, &n, *n)
}

// 2.4 pointer
func Pointer() {
	fmt.Println("\n2.34 pointer")

	var src = 2022
	var ptr = &src
	increase(ptr)
	fmt.Printf("调用increase(src)之后, n=%v,内存地址为%v,指向的原始值为%v\n", ptr, &ptr, *ptr)

	var ptr1 = new(int)
	fmt.Printf("ptr1=%v,内存地址为%v,指向的原始值为%v\n", ptr1, &ptr1, *ptr1)

}

// 2.5 fmt格式字符
func FmtVerbs() {
	fmt.Println("\n2.5 fmt格式字符")
	fmt.Println("\n2.5.1 通用")
	fmt.Printf("%%\n")

	fmt.Println("\n2.5.2 整数")
	i := 123
	fmt.Printf("%U\n", i)
	fmt.Printf("%c\n", i)
	fmt.Printf("%q\n", i)

	fmt.Println("\n2.5.3 浮点数")
	f := 123.456
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%20f\n", f)
	fmt.Printf("%b\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%X\n", f)

	fmt.Println("\n2.5.4 布尔类型")
	fmt.Printf("%t\n", f == 123.456)

	fmt.Println("\n2.5.5 字符串或byte切片")
	s := "hello world"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)

	fmt.Println("\n2.5.6 指针")
	p := &s
	fmt.Printf("%p\n", p)

}

// 2.6 运算符
func Operator() {
	fmt.Println("\n2.6.1 算数运算符")
	fmt.Printf("8%%3=%d\n", 8%3)

	i := 123
	i++
	fmt.Printf("i=%d", i)

	fmt.Println("\n2.6.2 位运算符")
	var b uint8 = 0b00111100
	fmt.Printf("b>>2=%b\n", b>>2)
	fmt.Printf("b<<2=%b\n", b<<2)

	var b1 uint8 = 0b00111100
	var b2 uint8 = 0b11001111
	fmt.Printf("b1&b2=%b\n", b1&b2)
	fmt.Printf("b1|b2=%b\n", b1|b2)
	fmt.Printf("b1^b2=%b\n", b1^b2)

	fmt.Println("\n2.6.3 赋值运算符")
	b += 3
	fmt.Printf("b=%d\n", b)

	fmt.Println("\n2.6.4 关系运算符")
	fmt.Printf("b1==b2?%t\n", b1 == b2)

	fmt.Println("\n2.6.5 逻辑运算符")
	bool1 := true
	bool2 := false
	fmt.Printf("bool1&&bool2=%t\n", bool1 && bool2)
	fmt.Printf("bool1||bool2=%t\n", bool1 || bool2)
	fmt.Printf("!bool1=%t\n", !bool1)

}

// 3.1 if...else
func IfElse() {
	fmt.Println()
	fmt.Println("3.1 if...else")

	fmt.Println("请输入你的年龄")
	var age int8
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("小朋友不要学编程哦")
	} else if age < 25 {
		fmt.Println("大朋友不要学编程哦")
	} else {
		fmt.Println("老朋友不要学编程哦")
	}
	fmt.Println("请输入你的年龄")

	fmt.Println()
	fmt.Println("3.1.1 if简短语句 ")
	if i := 3; i > 0 {
		fmt.Println("i=", i)
	} else if i > 3 {
		fmt.Println("i=", i)
	}
}

// 3.2 switch...case
func SwitchCase() {
	fmt.Println()
	fmt.Println("3.2 switch...case")
	fmt.Println("请输入你的年龄")
	var age int8
	fmt.Scanln(&age)
	switch {
	case age < 13:
		fmt.Println("小朋友不要学编程哦")
		// 继续向下匹配下一项
		fallthrough
	case age < 25:
		fmt.Println("大朋友不要学编程哦")
		// case 后面自动有 break
	default: // default 可以省略
		fmt.Println("老朋友不要学编程哦")
	}
}

// 3.3 for 循环
func ForLoop() {
	fmt.Println()
	fmt.Println("\n3.3.1 无限循环")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 11 {
			fmt.Println()
			break
		}
	}
	fmt.Println("\n3.3.2 条件循环")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()

	fmt.Println("\n3.3.3 标准for 循环")
	for j := 1; j < 11; j++ {
		fmt.Print(j, "\t")
	}
	fmt.Println()
}

// 3.4 label 和 goto
func LabelAndGoto() {
	fmt.Println()
	fmt.Println("\n3.4 label 和 goto")
	fmt.Println("\n3.4.1 label")

outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+ ")
			if i == 9 && j == 4 {
				fmt.Println()
				break outside
			}
		}
		fmt.Println()
	}

	fmt.Println("\n3.4.2 goto ")
	fmt.Println("1")
	if i := 1; i == 1 {
		goto four // 不推荐使用
	}
	fmt.Println("2")
	fmt.Println("3")
four:
	fmt.Println("4")
	fmt.Println("5")
}

// 3.5 函数
func getRes1(n1, n2 int) (int, int) {
	sum := n1 + n2
	diff := n1 - n2
	return sum, diff
}

func getRes(n1, n2 int) (sum, diff int) {
	sum = n1 + n2
	diff = n1 - n2
	return
}

func Function() {
	fmt.Println()
	fmt.Println("\n3.5 函数")
	res1, res2 := getRes(1, 2)
	fmt.Println("res1 =", res1, ",res2 =", res2)
	// 函数 != 调用函数, 饭 != 吃饭
	// fmt.Printf("getRes =%v,type of getRes =%T\n", getRes, getRes)
	// 匿名函数
	// var getRes2 = func(n1, n2 int) (sum, diff int) {
	// 	sum = n1 + n2
	// 	diff = n1 - n2
	// 	return
	// }
	// fmt.Printf("getRes =%v,type of getRes =%T\n", getRes2, getRes2)

	// 立即调用的匿名函数
	res3, res4 := func(n1, n2 int) (sum, diff int) {
		sum = n1 + n2
		diff = n1 - n2
		return
	}(3, 4)
	fmt.Println("res3 =", res3, ",res4 =", res4)

}

// 3.6 defer
func deferUtil() func(n int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被第%v次调用\n", i)
		return i
	}
}

func Defer() int {
	fmt.Println()
	fmt.Println("\n3.6 defer")
	fmt.Println("\n3.6.1 defer")
	f := deferUtil()
	// f(1)
	// defer f(f(3)) // f(3) 会被立即执行，f(f(3)) 会被延迟执行
	defer f(1)
	defer f(2)
	defer f(3)
	return f(4) // 顺序 4 -> 3 -> 2 -> 1
}

func DeferRecover() {
	// 不会中断程序执行，错误捕捉
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

// 全局变量首先初始化，然后是 init 函数
// 被依赖包的全局变量 >> 被依赖包的 init 函数 >> ... >> main 包的全局变量 >> main 的 init 函数 >> main 函数
var A = util.F("note.A")

// init 函数可以有多个，可以重名
func init() {
	util.F("note.init1")
}

func init() {
	util.F("note.init2")
}

// 4.1 数组
func Array() {
	fmt.Println()
	fmt.Println("\n4.1 数组")
	// var a [3]int = [3]int{1,
	// 	456,
	// 	789,
	// }
	var a = [...]int{1,
		456,
		789,
	}
	a[0] = 123
	fmt.Println("\n4.1.1 数组 for 遍历")
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v]=%v\n", i, a[i])
	}

	fmt.Println("\n4.1.2 for_range 遍历")
	for i, v := range a {
		fmt.Printf("a[%v]=%v\n", i, v)
	}
	fmt.Println("\n4.1.3 二维数组for_range 遍历")

	var twoDimensionalArray [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}
	for i, v := range twoDimensionalArray {
		for i2, v2 := range v {
			fmt.Printf("a[%v][%v]=%v\t", i, i2, v2)
		}
		fmt.Println()
	}
}

// 4.2 切片
func Slice() {
	fmt.Println()
	fmt.Println("\n4.2 切片")
	array := [5]int{1, 2, 3, 4, 5}
	// var s1 []int = array[0:len(array)]
	// array[0:len(array)] 等效于 array[:]
	var s1 []int = array[1:4] // [开始引用的index:结束引用的index+1)
	s1[0] = 0                 // array[1] 也会被改变，因为切片是对数组的引用，本身并不存储任何数据
	fmt.Println("array=", array)
	s2 := s1[1:]
	s2[0] = 0 // array 同样也会发生变化
	fmt.Println("array=", array)

	var s3 []int // 默认值是 nil
	fmt.Println("s3 =", s3 == nil)
	s3 = make([]int, 3, 5) // make([]Type,len,cap)底层创建一个长度为3的数组，最多可以存放5个，如果不写容量，默认 cap = len
	fmt.Println("s3 =", s3)
	fmt.Println("len(s3) =", len(s3))
	fmt.Println("cap(s3) =", cap(s3))

	s4 := []int{1, 2, 3} // 由系统自动创建底层数组
	fmt.Println("s4 =", s4)

	fmt.Println("\n4.2.3 追加元素")
	// 追加单个元素
	s1 = append(s1, 6, 7, 8) // 底层创建了新的数组，不再引用原数组
	s1[4] = 0
	fmt.Println("array=", array)
	fmt.Println("s1=", s1)
	// 追加其他切片
	s5 := append(s1, s2...)
	fmt.Println("s5=", s5)

	fmt.Println("\n4.2.4 复制数组")
	s6 := []int{1, 1}
	copy(s6, s5) // 容量能接收多少就接收多少
	fmt.Println("s6=", s6)

	fmt.Println("\n4.2.5 string与[]type")
	str := "hello 世界"
	fmt.Printf("[]byte(str)= %v\n[]byte(str)= %s\n", []byte(str), []byte(str))
	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	}

	// key := util.SelectByKey("注册", "登录", "注销")
	// fmt.Println("接收到 key =", key)
}

// 4.3 map
func Map() {
	fmt.Println("\n4.3 Map")
	var m1 map[string]string
	fmt.Println(m1 == nil)
	m1 = make(map[string]string, 2) // make(Type,初始size)
	m1["早上"] = "敲代码"
	m1["中午"] = "送外卖"
	m1["晚上"] = "开滴滴"
	fmt.Println("m1 = ", m1)
	m2 := map[string]string{
		"下午": "改bug",
		"凌晨": "卖早餐",
	}
	fmt.Println("m2 = ", m2)
	v, ok := m2["早上"]
	if ok {
		fmt.Println("v = ", v)
	} else {
		fmt.Println("key不存在, ok =", ok)
	}
	delete(m1, "晚上")
	fmt.Println("m1 = ", m1)
	// m1 = nil
	m2 = make(map[string]string)
	fmt.Println("m2 = ", m2)
	for key, value := range m1 {
		fmt.Printf("m1[%v] = %v\n", key, value)
	}
}

// 4.4 自定义数据类型&类型别名
func TypeDefinationAndTypeAlias() {
	fmt.Println()
	fmt.Println("\n4.4 自定义数据类型&类型别名")
	fmt.Println("\n4.4.1 自定义数据类型")
	type mesType uint16
	var textMes mesType = mesType(10000)
	fmt.Printf("textMs = %v, Type of textMes = %T \n", textMes, textMes)
	fmt.Println("\n4.4.2 类型别名")
	type myUnit16 = uint16
	var myu16 myUnit16 = 1000
	fmt.Printf("myUnit16 = %v, Type of myUnit16 = %T \n", myu16, myu16)
}

// 4.5 结构体
type User struct {
	Name string `json:"name"`
	Id   uint32
}

type Account struct {
	User
	password string
}

type Contact struct {
	*User
	Remark string
}

func Struct() {
	var u1 User = User{
		Name: "张三",
	}
	u1.Id = 10000
	// var u2 *User = new(User)
	var u2 *User = &User{
		Name: "李四",
	}
	// (*u2).Id = 10001
	u2.Id = 10001 // (*u2).Id = 10001 的简写形式
	var a1 = Account{
		User: User{
			Name: "王五",
		},
		password: "666",
	}
	var c1 *Contact = &Contact{
		User: &User{
			Id: u2.Id,
		},
		Remark: "王麻子",
	}
	c1.Name = "王五" // Name 字段唯一时，等效于 c1.User.Name = "王五"
	fmt.Println("a1 = ", a1)
	fmt.Println("c1 = ", c1)
	fmt.Println("c1.User = ", *((*c1).User))
}

// 5.1 方法
// 与特定类型绑定的函数，类型的定义和方法需要在同一个包内
// func(接收参数名 类型)方法名(形参列表)返回值列表{}
// 接收参数可以使用指针
func (u User) printName() {
	fmt.Println("u.Name =", u.Name)
}

func (u *User) setId() {
	(*u).Id = 10000
}

func Method() {
	u := User{
		Name: "小方块",
	}
	fmt.Println("\n5.1 方法")
	u.printName()
	u.setId()
	fmt.Println("u = ", u)
}

// 5.2 接口
type textMes struct {
	Type string
	Text string
}

// 因为是值传递，所以使用引用类型 *textMes，否则就是值拷贝了
func (tm *textMes) setText() {
	tm.Text = "hello"
}

type imgMes struct {
	Type string
	Img  string
}

func (im *imgMes) setImg() {
	im.Img = "清明上河图"
}

type Mes interface {
	setType()
}

func (tm *textMes) setType() {
	tm.Type = "文字消息"
}

func (im *imgMes) setType() {
	im.Type = "图片消息"
}

func SendMes(m Mes) {
	m.setType()
	switch mptr := m.(type) {
	case *textMes:
		mptr.setText()
	case *imgMes:
		mptr.setImg()
	}
	fmt.Println("m = ", m)
}

func Interface() {
	tm := textMes{}
	SendMes(&tm)
	im := imgMes{}
	SendMes(&im)
	var n1 int = 1
	// 类型断言
	n1interface := interface{}(n1)
	n2, ok := n1interface.(int)
	if ok {
		fmt.Println("n2 =", n2)
	} else {
		fmt.Println("类型断言失败")
	}
}

// 5.3 协程 Goroutine
var (
	c    int
	lock sync.Mutex
)

func PrimeNum(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	// fmt.Printf("%v\t", n)
	lock.Lock()
	c++
	lock.Unlock()
}

func Goroutine() {
	for i := 2; i < 100001; i++ {
		// 开启协程
		go PrimeNum(i)
	}
	fmt.Println()
	fmt.Println("\n5.2 协程 Goroutine")
	fmt.Printf("\n共找到%v个质数\n", c)
	var key string
	fmt.Scan(&key)
}

// 5.4 channel
func pushNum(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func pushPrimeNum(n int, c chan int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	c <- n
}

func Channel() {
	var c1 chan int = make(chan int)
	go pushNum(c1)
	// for {
	// 	v, ok := <-c1
	// 	if ok {
	// 		fmt.Printf("%v\t", v)
	// 	} else {
	// 		break
	// 	}
	// }
	for v := range c1 {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	var c2 chan int = make(chan int)
	for i := 2; i < 10001; i++ {
		go pushPrimeNum(i, c2)
	}
Print:
	for {
		select {
		case v := <-c2:
			fmt.Printf("%v\t", v)
		default:
			fmt.Println("\n所有质数都已找到")
			break Print
		}
	}
}
