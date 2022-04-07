[TOC]

# 基础数据类型

## 复数
- `complex(1,2)`: 1+2i

## iota
```go
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    ...
)
```
- 周日将对应0，周一为1，如此等等


