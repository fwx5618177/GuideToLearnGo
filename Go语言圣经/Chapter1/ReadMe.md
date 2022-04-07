[TOC]

# 基本结构

## 包管理
- 初始化：`go mod init [name]`

## 1. Hello World
```go
package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello World!");
}
```

## 2. 运行和编译
- Go是一门编译型语言，Go语言的工具链将源代码及其依赖转换成计算机的机器指令（译注：静态编译）
- 编译一个或多个以.go结尾的源文件，链接库文件，并运行最终生成的可执行文件: `go run [file]`
- 保存编译结果: `go build [file]`

## 3. 包
- Go语言的代码通过包（package）组织，包类似于其它语言里的库（libraries）或者模块（modules）
- 一个包由位于单个目录下的一个或多个.go源代码文件组成, 目录定义包的作用
- 每个源文件都以一条package声明语句开始，这个例子里就是package main, 表示该文件属于哪个包，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句

## 3. command line
- os: 以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
- os.Args变量是一个字符串（string）的切片（slice）（译注：slice和Python语言中的切片类似，是一个简版的动态数组）
- 左闭右开: a = [1, 2, 3, 4, 5], a[0:3] = [1, 2, 3]
- 0: 是命令本身的名字
- other: paramters

### 切片技巧
- form m to n: [m:n]
- [0, len(os.Args)]
- os.Args[0:]

```go
package main

import (
	"fmt"
	"os"
)

func main()  {
	var s, sep string

	for i := 0 ; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("Hello World!", s)
}
```

### for
- range: 产生一对值[index, value]
- _: 空标识符

```go
for _, arg := range os.Args[1:] {}
```

### 变量声明
```go
// 1.
s := ""

// 2.
var s string

// 3.
var s = ""

// 4.
var s string = ""
```
1. 简洁，只适合在函数内部,不能用于包变量
2. 依赖于字符串的默认初始化零值机制，被初始化为""
3. 同时声明多个变量
4. 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了


### strings.Join()
- 代替`+=`连接
- 使用`string`包的`Join`函数
```go
fmt.Println(strings.Join(os.Args[1:], " "));
```

# 习题
练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
```go
fmt.Println(os.Args[0])
```

练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
```go
    for index, args := range os.Args[0:] {
        fmt.Println(index, args)
    }
```

练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
```go
    start := time.Now()
    for index, args := range os.Args[0:] {
        fmt.Println(index, args)
    }
    end := time.Now()

    fmt.Println("end:", start, end)
```

## 4. 查找重复行

```go
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}
```

- 等价:
```go
counts[input.Text()]++

line := input.Text()
counts[line] = counts[line] + 1
```


### make - map
- 存储k-v集合，提供常数时间的存、取
- 用`==`比较
- 内置函数`make`创建空`map`
- go的map相当于: java-HashMap, Python-dict, Lua-table
- map就是映射

### bufio
- `Scan()`标准输入中读取内容

### Printf
```shell
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```

### os.Open
- 打开文件
- 返回两个值:
    1. 被打开的文件
    2. 错误
- err判断是否存在，则是`err != nil`, nil就是NULL

### strings.Split
- 字符串分割成子串的切片

# 网络客户端 - net
- `net/http`

# 并发-ch
```go
ch := make(chan string)
ch <- fmt.Sprint(err) // send to channel ch
fmt.Println(<-ch>) // receive from channel ch
```

# web
- 简单web服务器:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}

```


# 总结
1. 基础语法
2. 命名类型
3. 指针
4. 方法和接口
5. 包

