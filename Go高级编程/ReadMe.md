[TOC]

# 难点
- 并发
- CGO
- Go汇编
- RPC
- Protobuf
    - grpc
    - pbgo

# Go语言高级编程
- 目前阅读: [Go语言高级编程](https://books.studygolang.com/advanced-go-programming-book/)

# 数组
- 长度为0的数组在内存中并不占用空间。空数组虽然很少直接使用，但是可以用于强调某种特有类型的操作时避免分配额外的内存空间，比如用于管道的同步操作：
```go
c1 := make(chan [0]int)
go func() {
    fmt.Println("c1")
    c1 <- [0]int{}
}()
<-c1
```
- 我们并不关心管道中传输数据的真实类型，其中管道接收和发送操作只是用于消息的同步。对于这种场景，我们用空数组来作为管道类型可以减少管道元素赋值时的开销。当然一般更倾向于用无类型的匿名结构体代替：
```go
    c2 := make(chan struct{})
    go func() {
        fmt.Println("c2")
        c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
    }()
    <-c2
```

# 避免切片内存泄漏
切片操作并不会复制底层的数据。底层的数组会被保存在内存中，直到它不再被引用。但是有时候可能会因为一个小的内存引用而导致底层整个数组处于被使用的状态，这会延迟自动内存回收器对底层数组的回收

- FindPhoneNumber函数加载整个文件到内存，然后搜索第一个出现的电话号码，最后结果以切片方式返回
```go
func FindPhoneNumber(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return regexp.MustCompile("[0-9]+").Find(b)
}
```
- 这段代码返回的[]byte指向保存整个文件的数组。因为切片引用了整个原始数组，导致自动垃圾回收器不能及时释放底层数组的空间。一个小的需求可能导致需要长时间保存整个文件数据。这虽然这并不是传统意义上的内存泄漏，但是可能会拖慢系统的整体性能

- 要修复这个问题，可以将感兴趣的数据复制到一个新的切片中：
```go
func FindPhoneNumber(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = regexp.MustCompile("[0-9]+").Find(b)
    return append([]byte{}, b...)
}
```

# 接口-可变参数
当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：
```go
func main() {
    var a = []interface{}{123, "abc"}

    Print(a...) // 123 abc
    Print(a)    // [123 abc]
}

func Print(a ...interface{}) {
    fmt.Println(a...)
}
```

- 其中defer语句延迟执行了一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量v，这种函数我们一般叫闭包。闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问
- 闭包的这种引用方式访问外部变量的行为可能会导致一些隐含的问题：
```go
func main() {
    for i := 0; i < 3; i++ {
        defer func(){ println(i) } ()
    }
}
// Output:
// 3
// 3
// 3
```

- 因为是闭包，在for迭代语句中，每个defer语句延迟执行的函数引用的都是同一个i迭代变量，在循环结束后这个变量的值为3，因此最终输出的都是3


- 修复的思路是在每轮迭代中为每个defer函数生成独有的变量。可以用下面两种方式：
```go
func main() {
    for i := 0; i < 3; i++ {
        i := i // 定义一个循环体内局部变量i
        defer func(){ println(i) } ()
    }
}

func main() {
    for i := 0; i < 3; i++ {
        // 通过函数传入i
        // defer 语句会马上对调用参数求值
        defer func(i int){ println(i) } (i)
    }
}
```

# 原子操作
- 所谓的原子操作就是并发编程中“最小的且不可并行化”的操作。通常，如果多个并发体对同一个共享资源进行的操作是原子的话，那么同一时刻最多只能有一个并发体对该资源进行操作
- 一般情况下，原子操作都是通过“互斥”访问来保证的，通常由特殊的CPU指令提供保护

# 常见并发模式
- CSP：通讯顺序进程
- 并发编程的核心: 同步通信
- 并发更关注的是程序的设计层面，并发的程序完全是可以顺序执行的，只有在真正的多核CPU上才可能真正地同时运行。并行更关注的是程序的运行层面，并行一般是简单的大量重复，例如GPU中对图像处理都会有大量的并行运算
- 在并发编程中，对共享资源的正确访问需要精确的控制，在目前的绝大多数语言中，都是通过加锁等线程同步方案来解决这一困难问题，而Go语言却另辟蹊径，它将共享的值通过Channel传递(实际上多个独立执行的线程很少主动共享资源)

## 实现同步通信
1. sync.Mutex互斥锁同步
    - Lock()
    - Unlock()
2. 无缓存的管道: `<-`
3. 带缓存的管道
4. 等待多个线程完成后进行下一步同步: `sync.WaitGroup / wg.Add() / wg.Done() / wg.Wait()`

## 常见模型
- 生产者消费者
- 发布订阅模式

# CGO
- go build会在编译和链接阶段启动`gcc`
```go
package main

import "C"

func main() {
	println("hello cgo")
}
```




# 汇编
- Go汇编不是独立的语言
- 查看go对应的伪汇编代码: `go tool compile -S 【file】`
    - S:输出汇编格式
- [go 汇编入门指南](https://go.dev/doc/asm)
- 如果是纯粹学习汇编语言，则可以从《深入理解程序设计：使用Linux汇编语言》开始，该书讲述了如何以C语言的思维变现汇编程序。如果是学习X86汇编，则可以从《汇编语言：基于x86处理器》一开始，然后再结合《现代x86汇编语言程序设计》学习AVX等高级汇编指令的使用。

Go汇编语言的官方文档非常匮乏。其中“A Quick Guide to Go's Assembler”是唯一的一篇系统讲述Go汇编语言的官方文章，该文章中又引入了另外两篇Plan9的文档：A Manual for the Plan 9 assembler 和 Plan 9 C Compilers。Plan9的两篇文档分别讲述了汇编语言以及和汇编有关联的C语言编译器的细节。看过这几篇文档之后会对Go汇编语言有了一些模糊的概念，剩下的就是在实战中通过代码学习了。


# RPC
- 基于`net/rpc`的扩展
- 规范设计
    - client - 规范 - server
- 封装内部方法
- 跨语言


# go web
- 框架：
    - router：httpRouter
    - MVC：beego，goa，chi

# 分布式系统
## 分布式id生成器
- snowflake
- sonyflake
## 分布式锁
- trylock: goroutine在抢锁失败后，需要放弃其流程
- setnx:基于Redis,抢占逻辑,很适合在高并发场景下，用来争抢一些“唯一”的资源
- zookeeper: Lock成功之前会一直阻塞,分布式的阻塞锁比较适合分布式任务调度场景，但不适合高频次持锁时间短的抢锁场景
- etcd

## 延时任务系统
- 定时器
    - 时间堆：小顶堆，特殊二叉树
    - 时间轮：哈希表
- 定时任务触发：
    - 将任务被触发的信息封装为一条消息，发往消息队列，由用户侧对消息队列进行监听。
    - 对用户预先配置的回调函数进行调用。
## 分布式搜索引擎
- elasticsearch

## 负载均衡
- 洗牌算法
- 修正洗牌算法(fisher-yates)

