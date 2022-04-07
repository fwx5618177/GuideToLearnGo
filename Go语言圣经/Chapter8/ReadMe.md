# CSP
- Go语言中的并发程序可以用两种手段来实现。本章讲解goroutine和channel，其支持“顺序通信进程”(communicating sequential processes)或被简称为CSP
- CSP是一种现代的并发编程模型，在这种编程模型中值会在不同的运行实例(goroutine)中传递
- 传统的并发模型：多线程共享内存

# Goroutines
- 在Go语言中，每一个并发的执行单元叫作一个goroutine
- 当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行

```go
f() // wait for call f() to return

go f() // create a new goroutine that calls f(); don't wait
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)

	const n = 45

	fibN := fib(n)
	fmt.Printf("%d, %d", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)

			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}


```

- 主函数返回时，所有的goroutine都会被直接打断，程序退出

# 并发的Clock服务

```go
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}


	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)

			continue
		}

		handleConn(conn)
	}

}


func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
```

## defer
- 可以使用关键字defer向函数注册退出调用，即函数退出时，defer后的函数才被调用
- defer语句的作用是不管程序是否出现异常，均在函数退出时自动执行相关代码

# channels
- 如果说goroutine是Go语言程序的并发体的话，那么channels则是它们之间的通信机制
- 一个channel是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息
- 每个channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int
- 一个channel有发送和接受两个主要操作，都是通信行为

```go
ch := make(chan int)


ch <- x  // a send statement
x = <-ch // a receive expression in an assignment statement
<-ch     // a receive statement; result is discarded

close(ch)
```

# make
- 最简单方式调用make函数创建的是一个无缓存的channel，但是我们也可以指定第二个整型参数，对应channel的容量
- 如果channel的容量大于零，那么该channel就是带缓存的channel
```go
ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

# 不带缓存的Channels
- 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句
- 如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作
- 基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原因，无缓存Channels有时候也被称为同步Channels





# 串联的Channels（Pipeline）

- Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入。这种串联的Channels就是所谓的管道（pipeline）
- 像这样的串联Channels的管道（Pipelines）可以用在需要长时间运行的服务中，每个长时间运行的goroutine可能会包含一个死循环，在不同goroutine的死循环内部使用串联的Channels来通信

```go
package main

import (
	"fmt"
)

func main() {
	na := make(chan int)
	sq := make(chan int)

	go func () {
		for x := 0; ;x ++ {
			na <- x
		}
	}()

	go func() {
		for {
			x := <- na

			sq <- x * x
		}
	}()

	for {
		fmt.Println(<-sq)
	}
}

```

- 当一个channel被关闭后，再向该channel发送数据将导致panic异常。当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值。关闭上面例子中的naturals变量对应的channel并不能终止循环，它依然会收到一个永无休止的零值序列，然后将它们发送给打印者goroutine
- 使用`close(sq)`

- 先关闭na
- 在关闭sq
- range可以在channels上迭代

```go
package main

import (
	"fmt"
)

func main() {
	na := make(chan int)
	sq := make(chan int)

	go func () {
		for x := 0; x < 100 ;x ++ {
			na <- x
		}

		close(na)
	}()

	go func() {
		for x := range na {
			sq <- x * x
		}

		close(sq)
	}()

	for x := range sq {
		fmt.Println(x)
	}
	
}

```

# 单方向的channels
- 当一个channel作为一个函数参数时，它一般总是被专门用于只发送或者只接收
- Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel
- 类型chan<- int表示一个只发送int的channel，只能发送不能接收
- 类型<-chan int表示一个只接收int的channel，只能接收不能发送

```go
package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}

	close(out)
}

func square(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}

	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	na := make(chan int)
	sq := make(chan int)

	go counter(na)
	go square(sq, na)
	printer(sq)
	
}

```


# 带缓存的Channels
- 带缓存的Channel内部持有一个元素队列
- `ch = make(chan string, 3)`
- 向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素
- 如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。

## cap
- 获取channel内部缓存的容量： `cap(ch)`
- 返回channel内部缓存队列中有效元素的个数： `len(ch)`
- Go语言新手有时候会将一个带缓存的channel当作同一个goroutine中的队列使用，虽然语法看似简单，但实际上这是一个错误。Channel和goroutine的调度器机制是紧密相连的，如果没有其他goroutine从channel接收，发送者——或许是整个程序——将会面临永远阻塞的风险。如果你只是需要一个简单的队列，使用slice就可以了。
- 多个goroutines并发地向同一个channel发送数据，或从同一个channel接收数据都是常见的用法
```go
func mirroredQuery() string {
	response := make(chan string, 3)

	go func() { response <- request("1") }()
	go func() { response <- request("2") }()
	go func() { response <- request("3") }()

	return <- response
}

func request(hostname string) (response string) {
	
	return ""
}
```

# 并发的循环
- 指定buffer的channel

```go
package main

import (
	"fmt"
)

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfiles string
		err error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item

			it.thumbfiles, it.err = thumbfiles.ImageFile(f)

			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch

		if it.err != nil {
			return nil, it.err
		}

		thumbfiles = append(thumbfiles, it.thumbfiles)
	}

	return thumbfiles, nil
}

```

# 并行爬虫
- 可以用一个有容量限制的buffered channel来控制并发，这类似于操作系统里的计数信号量概念。从概念上讲，channel里的n个空槽代表n个可以处理内容的token(通行证)，从channel里接收一个值会释放其中的一个token，并且生成一个新的空槽位。这样保证了在没有接收介入时最多有n个发送操作
- [crawk](https://docs.hacknode.org/gopl-zh/ch8/ch8-06.html)

# select实现多路复用
```go
select {
case <-ch1:
    // ...
case x := <-ch2:
    // ...use x...
case ch3 <- y:
    // ...
default:
    // ...
}
```







