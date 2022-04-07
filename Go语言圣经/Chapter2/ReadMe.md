[TOC]
# 程序结构

## 保留字
```shell
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var

内建常量: true false iota nil

内建类型: int int8 int16 int32 int64
          uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error

内建函数: make len cap new append copy close delete
          complex real imag
          panic recover
```

## 声明
- var: 变量
- const: 常量
- type: 类型
- func: 函数

### var
```go
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four"
```

### := 
`i := 100`

### 指针
```go
x := 1
p := &x // p, of type *int, points to x
*p = 2 // equivalent to x = 2
```

### new
```go
p := new(int) // p, *int
```

## 生命周期
- 变量的生命周期指的是在程序运行期间变量有效存在的时间段
- 对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的
- 而相比之下，局部变量的生命周期则是动态的：
    - 每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收
    - 函数的参数变量和返回值变量都是局部变量。它们在函数每次被调用的时候创建。

# 导入包
- 一个导入路径代表一个目录中的一个或多个Go源文件
- `gopl.io/ch2/tempconv`包的名字一般是`tempconv`
- 包内的成员将通过类似tempconv.CToF的形式访问
- 包级别的名字，在一个文件声明的类型和常量，在同一个包的其他源文件也是可以直接访问的
