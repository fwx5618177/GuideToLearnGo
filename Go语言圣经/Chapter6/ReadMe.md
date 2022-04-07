# 方法
- OOP

# 方法声明
- 在函数声明时，在其名字之前放上一个变量，即是一个方法
- 这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法
```go
type Point struct{x, y float64}

// traditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
    return 1.111
}
```

- p: 方法的接收器
- 早期的面向对象语言留下的遗产将调用一个方法称为“向一个对象发送消息”
- 在Go语言中，我们并不会像其它语言那样用this或者self作为接收器；我们可以任意的选择接收器的名字。由于接收器的名字经常会被使用到，所以保持其在方法间传递时的一致性和简短性是不错的主意。这里的建议是可以使用其类型的第一个字母，比如这里使用了Point的首字母p
```go
p := Point{1,2}
q := Point{4,5}

Distance(p, q) // 5, function call
p.Distance(q)   // 5, method call
```

- 除了`stcurt`，也可以定义方法
```go
type Path []Point

func (path Path) Distance(param int) float64 {}

perim := Path {
    {1,2},
    {3,4},
}

perim.Distance(1)
```


# 基于指针对象的方法
- 当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需要用到指针了
- 当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法

```go
func (p *Point) ScaleBy(factor float64) float64 {}
```
- 这个方法的名字是(*Point).ScaleBy
- 这里的括号是必须的,没有括号的话这个表达式可能会被理解为*(Point.ScaleBy)

- 在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的
- 只有类型(Point)和指向他们的指针(*Point)，才可能是出现在接收器声明里的两种接收器
```go
type P *int
func (P) f() { /* ... */ } // compile error: invalid receiver type
```

- 想要调用指针类型方法(*Point).ScaleBy，只要提供一个Point类型的指针即可:
1.
```go
r := &Point{1, 2}
r.ScaleBy(2)
fmt.Println(*r) // "{2, 4}"
```
2. 
```go
p := Point{1, 2}
pptr := &p
pptr.ScaleBy(2)
```
3.
```go
p := Point{1, 2}
(&p).ScaleBy(2)
fmt.Println(p) // "{2, 4}"
```

# callback
```go
type Rocket struct {}

func (r *Rocket) Launch() {}

r := new(Rocket)

x.funcA(r.Launch)
```

# Bit数组
- 在数据流分析领域，集合元素通常是一个非负整数，集合会包含很多元素，并且集合会经常进行并集、交集操作，这种情况下，bit数组会比map表现更加理想。
- 比如我们执行一个http下载任务，把文件按照16kb一块划分为很多块，需要有一个全局变量来标识哪些块下载完成了，这种时候也需要用到bit数组
`words []unit64`


## String方法
- 类似于toString

# 封装
- 一个对象的变量或者方法如果对调用方是不可见的话，一般就被定义为“封装”，封装有时候也被叫做信息隐藏