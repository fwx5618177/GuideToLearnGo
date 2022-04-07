[TOC]

# 复合基础类型
- slice
- 数组
- map
- 结构体

## 数组声明
```go
q := [3]int{1,2,3}

type Currency int

const (
    USD Currency = iota
    EUR
    GBP
    RMB
)

symbol := [...]string{USD: "$", EUR: "$"}
```

## 切片
- `a[:]`
- `a[1:]`

## Map
- map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作
```go
ages := make(map[string]int)

// equavilent
age := map[string]int {
    "alice": 31,
    "cha": 34,
}
```

# 结构体
```go
type Employee struct {
    ID int
    Name string
    Address string
}

var di Employee
```
- 字面值
```go
type Point struct { X, Y int}

p := Point{1,2}
```

- 嵌入和匿名:
```go
type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}
```

## JSON
- 编组
```go
type Movie struct {
    title string
    Year int `json:"released"`
}

var movies = []Movie{
    {
        title: "A",
        Year: 1942,
    },
    {
        title: "B",
        Year: 1999,
    }
}

data, err := json.Marshal(movies)
```







