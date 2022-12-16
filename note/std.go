package note

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"gonote/util"
	"io"
	"math/rand"
	"net"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// 6.1 随机数
func Random() {
	fmt.Println()
	fmt.Println("\n6.1 随机数")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10) + 1)
}

// 6.2 字符串类型转换
func StrConv() {
	fmt.Println()
	fmt.Println("\n6.2 字符串类型转换")
	i1 := 123
	s1 := "lijingyuan.top"
	s2 := fmt.Sprintf("%d@%s", i1, s1)
	fmt.Println(s2)

	var (
		i2 int
		s3 string
	)
	n, err := fmt.Sscanf(s2, "%d@%s", &i2, &s3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n成功解析了%d个数据\n", n)
	fmt.Println("i2 =", i2)
	fmt.Println("s3 =", s3)

	s4 := strconv.FormatInt(123, 4)
	fmt.Println("s4 =", s4)
	u1, err := strconv.ParseInt(s4, 4, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("u1 =", u1)

}

// 6.3 字符串常见操作
func PackageStrings() {
	fmt.Println()
	fmt.Println("\n6.3 字符串常见操作")
	fmt.Println(strings.Contains("hello", "o"))
	fmt.Println(strings.Index("hello", "ll"))
	fmt.Println(strings.Replace("hello", "l", "dd", -1))
	fmt.Println(strings.Repeat("mia", 5))
	fmt.Println(strings.Fields("mia mia\n mia\tmia"))
	fmt.Println(strings.Split("he-llo-world", "-"))
	fmt.Println(strings.SplitAfter("he-llo-world", "-"))
	fmt.Println(strings.Trim("#*\nwww.baidu.com&^@", "#*\n&^@"))
}

// 6.4 中文字符常见操作
func PackageUtf8() {
	fmt.Println()
	fmt.Println("\n6.4 中文字符常见操作")
	str := "hello,世界"
	// 统计字符串切片符文字数
	fmt.Println(utf8.RuneCountInString(str))
	// 判断是否由有效的符文组成
	fmt.Println(utf8.ValidString(str[:len(str)-1]))
}

// 6.5 时间常见操作PackageTime
func PackageTime() {
	fmt.Println()
	fmt.Println("\n6.5 时间常见操作PackageTime")
	fmt.Println("\n6.5.1 时段")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println()
	d1, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Println("d1 =", d1)
	t1, err := time.Parse("2006年1月2日, 15点4分", "2023年1月1日, 18点18分")
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(t1))

	var intChan chan int = make(chan int)
	select {
	case <-intChan:
		fmt.Println("收到了用户发送的验证码")
	case <-time.After(1 * time.Second):
		fmt.Println("验证码已过期")
	}
	l1, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	fmt.Println(l1.String())

	fmt.Println()
	fmt.Println("\n6.5.5 时段")
	fmt.Println(time.Now().Format("2006年1月2日, 15点4分"))
	t2, err := time.ParseInLocation("2006年1月2日, 15点4分", "2122年12月13日, 21点39分", l1)
	if err != nil {
		panic(err)
	}
	fmt.Println(t2)
	fmt.Println(t2.Location())
	fmt.Println(t2.Add(d1))

	fmt.Println()
	fmt.Println("\n6.5.6 周期计数器")
	go func() {
		time.Sleep(1 * time.Second)
		intChan <- 1
	}()
TickerFor:
	for {
		select {
		case <-intChan:
			fmt.Println()
			break TickerFor
		case <-time.NewTicker(100 * time.Millisecond).C:
			fmt.Print(".")
		}
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("\n6.5.7 单次计数器")
	select {
	case <-intChan:
		fmt.Println("收到了用户发送的验证码")
	case <-time.NewTimer(time.Second).C:
		fmt.Println("验证码已过期")
	}
}

// 6.6 文件常见操作
func FileOperation() {
	fmt.Println()
	fmt.Println("\n6.6 文件常见操作")
	// util.MkdirWithFilePath("d1/d2/file2")

	fmt.Println("\n6.6.5 文件夹常见操作")
	dirEntrys, err := os.ReadDir("/Users/kangjinghang/logs/xxl-job")
	if err != nil {
		panic(err)
	}
	for _, v := range dirEntrys {
		fmt.Println(v.Name())
	}
	fmt.Println("\n6.6.6 文件常见操作")
	file, err := os.OpenFile("f1", os.O_RDWR|os.O_CREATE, 0665)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("\n6.6.7 无缓冲区读写(适合小文件)")
	data, err := os.ReadFile("f1")
	if err != nil {
		panic(err)
	}
	fmt.Println("f1中数据为: ", string(data))
	err = os.WriteFile("f2", data, 0775)
	if err != nil {
		panic(err)
	}
}

// 6.7 文件读写
func FileReadAndWrite() {
	fmt.Println()
	fmt.Println("\n6.7 文件读写")
	f5, err := os.OpenFile("f5", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f5.Close()
	writer := bufio.NewWriter(f5)
	for i := 1; i < 5; i++ {
		fileName := fmt.Sprintf("f%v", i)
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		data = append(data, '\n')
		writer.Write(data) // 写入缓冲区
	}
	writer.Flush()
}

// 6.8 错误
func Errors() {
	fmt.Println()
	fmt.Println("\n6.8 错误")
	// 捕捉错误
	defer func() {
		err := recover()
		fmt.Println("捕捉到了错误", err)
	}()
	err1 := errors.New("可爱的错误")
	fmt.Println("err1 =", err1)
	panic(err1) // 如果没有捕捉，之后的不会被运行
	err2 := fmt.Errorf("%s的错误", "温柔")
	fmt.Println("err2 =", err2)

}

// 6.9 日志
func Log() {
	fmt.Println()
	fmt.Println("\n6.9 日志")
	// 捕捉错误
	defer func() {
		err := recover()
		fmt.Println("捕捉到了错误", err)
	}()
	err := errors.New("可爱的错误")
	util.INFO.Print(err)
	util.WARN.Panic(err)
	// util.ERR.Fatal(err)
}

// 6.10 单元测试
func IsNotNegative(n int) bool {
	return n > -1
}

// 6.11 命令行参数
func CmdArgs() {
	fmt.Println()
	fmt.Println("\n6.11 命令行参数")
	fmt.Printf("接受到了%v个参数\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("第%v个参数是%v\n", i, v)
	}
	fmt.Println()
	vPtr := flag.Bool("v", false, "GoNote版本")
	var username string
	flag.StringVar(&username, "u", "", "用户名")
	flag.Func("f", "", func(s string) error {
		fmt.Println("执行函数, s =", s)
		return nil
	})
	flag.Parse()
	if *vPtr {
		fmt.Printf("GoNote版本是 v0.0.0\n")
	}
	fmt.Println()
	fmt.Println("当前用户为:", username)

	for i, v := range flag.Args() {
		fmt.Printf("第%v个无flag参数是%v\n", i, v)
	}
}

// 6.12 buildin包一览
func PackageBuildIn() {
	fmt.Println()
	fmt.Println("\n6.12 buildin包一览")
	c1 := complex(12.34, 45.67)
	fmt.Println("c1 =", c1)
	r1 := real(c1)
	i1 := imag(c1)
	fmt.Println("r1 =", r1)
	fmt.Println("i1 =", i1)
}

// 6.13 runtime包
func PackageRuntime() {
	if runtime.NumCPU() > 7 {
		runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	}
	// runtime.Goexit() // 终止当前协程，不影响 defer 语句
}

// 6.14 sync包
func PackageSync() {
	fmt.Println()
	fmt.Println("\n6.14 sync包")
	fmt.Println("\n6.14.1 Mutex互斥锁")
	fmt.Println("\n6.14.2 WaitGroup")

	var c int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	primeNum := func(n int) {
		defer wg.Done()
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return
			}
		}
		mutex.Lock()
		c++
		mutex.Unlock()
	}

	for i := 2; i < 100001; i++ {
		// 开启协程
		wg.Add(1)
		go primeNum(i)
	}
	wg.Wait()
	fmt.Printf("\n共找到%v个质数\n", c)

	fmt.Println("\n6.14.3 Cond")
	cond := sync.NewCond(&mutex)
	for i := 0; i < 10; i++ {
		go func(n int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Printf("协程%v被唤醒了\n", n)
			cond.L.Unlock()
		}(i)
	}
	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		if i == 4 {
			fmt.Println()
			cond.Signal() // 唤醒一个
		}
		if i == 9 {
			fmt.Println()
			cond.Broadcast() // 唤醒所有
		}
	}
	fmt.Println()
	fmt.Println("\n6.14.4 Once")
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("只有一次机会")
			})
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("\n6.14.5 Sync Map")
	var m sync.Map
	m.Store(1, 100)
	m.Store(2, 200)
	m.Store(3, 300)
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("m[%v]=%v\n", key, value.(int))
		return true
	})
}

// 7.4 sort包
type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (ps PersonSlice) Len() int {
	return len(ps)
}

func (ps PersonSlice) Less(i, j int) bool {
	return ps[i].Age > ps[j].Age
}

func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func PackageSort() {
	fmt.Println()
	fmt.Println("\n7.4 sort包")
	fmt.Println("\n7.4.1 对常见类型进行排序")
	is := []int{2, 4, 8, 10}
	v := 6
	i := sort.SearchInts(is, v)
	fmt.Printf("%v适合插入在%v的%v\n", v, is, i)

	fmt.Println("\n7.4.2 自定义排序")
	p := []Person{{"小小", 18}, {"小方", 5}, {"小块", 50}}
	sort.Slice(p, func(i, j int) bool {
		return p[i].Age < p[j].Age
	})
	fmt.Println("p =", p)

	fmt.Println("\n7.4.3 自定义查找")
	i = sort.Search(len(is), func(i int) bool {
		return is[i] >= v
	})
	fmt.Printf("%v中第一次出现不小于%v的位置是%v\n", is, v, i)

	fmt.Println("\n7.4.4 sort.Interface")
	sort.Sort(sort.Reverse(PersonSlice(p)))
	fmt.Println("p =", p)

}

// 7.5.2 二分查找
func binarySearch(s []int, key int) int {
	startIndex := 0
	endIndex := len(s) - 1
	midIndex := 0
	for startIndex <= endIndex {
		midIndex = (startIndex + endIndex) / 2
		if s[midIndex] < key {
			startIndex = midIndex + 1
		} else if s[midIndex] > key {
			endIndex = midIndex - 1
		} else {
			return midIndex
		}
	}
	return -1
}

func PackageBinarySearch() {
	fmt.Println("\n7.4.2 二分查找")
	s := make([]int, util.RandInt(20)+1)
	for i := 0; i < len(s); i++ {
		s[i] = util.RandInt(1000)
	}
	sort.Ints(s)
	q := 555
	i := binarySearch(s, q)
	if i == -1 {
		fmt.Printf("没有找到%v", q)
	} else {
		fmt.Printf("%v的下标是%v", q, i)
	}
}

// 9.1 JSON常见操作
func PackageJson() {
	fmt.Println()
	fmt.Println("\n9.1 JSON常见操作")
	type user struct {
		Name  string `json:"name"`
		Age   int    `json:"age,omitempty"` // 编码时，如果该字段为零值则跳过该字段
		Email string `json:"-"`             // 编码时，跳过该字段
		Job   map[string]string
	}
	u1 := user{
		Name: "方块",
		Age:  3,
		Job: map[string]string{
			"早班": "保安",
			"午班": "洗碗",
			"晚班": "送外卖",
		},
	}
	data, _ := json.Marshal(u1)
	fmt.Println(string(data))

	buf := new(bytes.Buffer)
	// json.Indent(buf, data, "+", "-=")
	json.Indent(buf, data, "", "\t")
	fmt.Println(buf.String())

	var u2 user
	json.Unmarshal(data, &u2)
	fmt.Println("u2 =", u2)
}

// 10.1 TCP编程入门
func TcpCli() {
	conn, err := net.Dial("tcp", "127.0.0.1:2022")
	if err != nil {
		fmt.Println("拨号失败")
		return
	}
	defer conn.Close()
	for {
		mes := struct {
			UserName string
			Mes      string
		}{
			UserName: "方块",
		}
		fmt.Println("请输入要发送的内容:")
		fmt.Scanf("%s\n", &mes.Mes)
		if mes.Mes == "" {
			fmt.Println("输入为空")
			continue
		}
		if mes.Mes == "exit" {
			return
		}
		// 复杂版:
		// data, _ := json.Marshal(&mes)
		// n, err := conn.Write(data)
		// if err != nil {
		// 	fmt.Println("发送失败")
		// 	return
		// }
		// fmt.Printf("成功发送了%v个字节\n", n)
		// 简单版:
		err := json.NewEncoder(conn).Encode(&mes)
		if err != nil {
			fmt.Println("发送失败")
			return
		}
	}
}

func TcpServer() {
	listener, err := net.Listen("tcp", ":2022")
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer listener.Close()
	for {
		fmt.Println("主进程等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接听失败:", err)
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("一个客户端协程已经开启...")
			defer conn.Close()
			for {
				// 复杂版:
				// buf := make([]byte, 4096)
				// n, err := conn.Read(buf)
				// if err == io.EOF {
				// 	fmt.Println("客户端退出了")
				// 	return
				// }
				// if err != nil {
				// 	fmt.Println("读取失败:", err)
				// 	return
				// }
				// mes := struct {
				// 	UserName string
				// 	Mes      string
				// }{}
				// json.Unmarshal(buf[:n], &mes)

				// 简化版:
				mes := struct {
					UserName string
					Mes      string
				}{}
				err := json.NewDecoder(conn).Decode(&mes)
				if err == io.EOF {
					fmt.Println("客户端退出了")
					return
				}
				if err != nil {
					fmt.Println("读取失败:", err)
					return
				}
				fmt.Printf("%s说:%s\n", mes.UserName, mes.Mes)
			}
		}(conn)
	}
}
