[TOC]

# Go 语言编程
- 单元测试
- 并发
- 游戏服务器编程
- 网络编程
    - RPC
    - Socket:TCP/UDP/ICMP
    - HTTP
    - JSON
- 加解密
    - 文件服务器
- 反射
- CGO
    

## 目录结构
```shell
├── clacproj
│   ├── bin
│   ├── pkg
│   └── src
│       ├── clac
│       └── simple
```

- go install: 在本目录下生成bin目录
    - go install 的编译过程有如下规律：
    - go install 是建立在 GOPATH 上的，无法在独立的目录里使用 go install。
    - GOPATH 下的 bin 目录放置的是使用 go install 生成的可执行文件，可执行文件的名称来自于编译时的包名。
    - go install 输出目录始终为 GOPATH 下的 bin 目录，无法使用-o附加参数进行自定义。
    GOPATH 下的 pkg 目录放置的是编译期间的中间文件。

- export GOPATH=[dir]

# 继承
```go
type Base struct {
	Name string
}

func (base *Base) Foo() {}
func (base *Base) Bar() {}

type Foo struct {
	Base
	age int
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
}
```

# 单元测试
```go
package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("new music manager failed.")
	}

	if mm.Len() != 0 {
		t.Error("not empty.")
	}

	m0 := &MusicEntry {
		"1", "Heart", "Pop", "1", "Mp3",
	}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("Add() failed.")
	}

	m, _ := mm.Find(m0.Name)

	if m == nil {
		t.Error("Find() failed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("found item failed.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("Get() failed.")

	}

	m, _ = mm.Remove(0)

	if m == nil || mm.Len() != 0 {
		t.Error("Remove() failed.", err)

	}
}


```

# 并发基础
模型：
1. 多进程
2. 多线程
3. 非阻塞/异步IO
4. 协程

并行通信模式：
1. 共享内存系统
2. 消息传递系统

## 协程
- 轻量级进程，可以创建几百万个，而进程最多只能不超过1万个

## 并发通信
1. 内存共享
C版本:
```c
#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>


pthread_mutex_t mutex1 = PTHREAD_MUTEX_INITIALIZER;
int counter = 0;

void *count();

int main(void) {
    int rc1, rc2;
    pthread_t thread1, thread2;

    /**
     * @brief 创建线程
     * 
     */
    if((rc1 = pthread_create(&thread1, NULL, &add, NULL))) {
        printf("Thread creation failed: %d\n", rc1);
    }

    if((rc2 = pthread_create(&thread2, NULL, &add, NULL))) {
        printf("Threadd creation failed: %d\n", rc2);
    }

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    exit(0);

    return 0;
}


void *count() {
    pthread_mutex_lock( &mutex1 );

    counter++;
    
    printf("Counter value: %d\n", counter);
    
    pthread_mutex_unlock( &mutex1 );
}
```

Go版本:
```go
package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()

	counter++

	fmt.Println(counter)

	lock.Unlock()
}

func main() {
	lock := &sync.Mutex()

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()

		if c >= 10 {
			break
		}
	}
}

```

缺点：
- 在一个大的系统中具有无数的锁、无数的共享变量、无数的业务逻辑与错误处理分
支

2. 消息机制
- 通过channel进行数据通信
- select()函数来监控一系列的文件句柄，一旦其中的一个文件句柄发生了IO动作，该select()调用就会被返回，常被用于实现高并发的Socket服务器程序
```go
select {
    case <- chan1:
    case chan2 <- 1:
    default:
}
```

## 超时机制
- 向channel写数据时发现channel已满
- 利用select解决超时问题

```go
timeout := make(chan bool, 1)

go func() {
	time.Sleep(1e9)
	timeout <- true
}()

select {
	case <- ch:
		// 从ch中获取数据
	case <- timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
}
```

# 多核并行化
- 本质是单线程的，还是没有多CPU运行
- 设置环境变量为: `runtime.GOMAXPROCS(16)`控制使用多少CPU核心
```go
package main

import (
	"fmt"
)

func Count(ch chan<- int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {

		go Count(chs[i])
	}
}

type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan<- int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}

	c <- 1
}

const NCPU = 16

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU)

	for i := 0; i < NCPU; i++ {

		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}

	for i := 0; i < NCPU; i++ {
		<- c
	}
}
```

# 出让时间片
- runtime包中的`Gosched()`：控制何时主动出让时间片给其他的goroutine

# RPC编程
- 一种通过网络从远程计算机程序上请求服
务，而不需要了解底层网络细节的应用程序通信协议
- 构建在TCP/UDP/HTTP上，开发者无需额外地为这个调用过程编写网络通信相关代码，使得开发包括网络分布式程序在内的应用程序更加容易
- `net/rpc`

# Gotool
- 文件内容格式化： go fmt [file]
- 获取远程包： go get [url]
- 文件格式:
 - 一个标准的Go语言工程包含以下几个目录：src、pkg和bin
- 自动生成文档: go doc [file]
- 工程构建会在当前所在目录生成可执行文件：`go build [file]`
- 安装到恰当的位置: `go install [file]`
- 环境变量: `export GOPATH=~/work/go-proj1:~/work2/goproj2:~/work3/work4/go-proj3`

# 跨平台开发
## 交叉编译
- 交叉编译是在一个平台上生成另一个平台上的可执行代码。 同一个体系结构可以运行不同的操作系统；同样，同一个操作系统也可以在不同的体系结构上运行。
- 如果你当前的编译目标为AMD64架构的64位Linux，那么Go包对应的安装位置是linux_amd64。推导之，如果当前的编译目标为x86架构的32位Windows，对应的安装位置就是windows_386
- 如果我们要在一台安装了64位Linux操作系统的AMD64电脑上执行一段Go代码，就必须用能够生成64位ELF文件格式的Go编译器进行编译和链接

- win: `GOOS=windows GOARCH=386 go build -o upload.exe upload.go`
- android: `GOOS=linux GOARCH=ARM `


