[TOC]

# 接口

- 接口类型是对其它类型行为的抽象和概括
- Go语言中接口类型的独特之处在于它是满足隐式实现的
- 在Go语言中还存在着另外一种类型：接口类型。接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么

```go
func Fprint(w io.Writer, format string, args ...interface{}) (int, error)

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

# 接口类型

- 接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例
```go
package io

type Reader interface {
    Read(p []byte) (n int, err error)
}

type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Reader
}
```
- 接口可以组合

# 接口的实现

```go
package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Length() float64
}

type Rect struct {
	width float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Length() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s Shape

	s = Rect{4.0, 5.0}
	
	fmt.Println("Area is:", s.Area())
}

```

# 类型断言
- x.(T)
- 一个类型断言检查它操作对象的动态类型是否和断言的类型匹配
