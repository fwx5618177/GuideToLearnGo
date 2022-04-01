[TOC]

# 函数

```go
func name(parameter) result {}
```

- 多返回值: `return msg, img, err`

# 函数值
```go

func square(n int) int { return n * n }

f := square

var g func(int) int

```

# 匿名函数
```go
func square() func() int {
    var x int

    return func() int {
        return 1
    }
}
```

# 可变参数
```go
func sum (vals...int) int {}
```

# defer
- 类似于js的`async-await`
- 直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束

```go

defer resp.Body.Close()

```


- 对文件操作:
```go
package ioutil

import (
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	return ReadAll(f)
}
```

- 处理互斥锁:
```go
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()

	defer mu.Unlock()

	return m[key]
}
```

# panic
- 当panic异常发生时，程序会中断运行，并立即执行在该goroutine（可以先理解成线程，在第8章会详细介绍）中被延迟的函数（defer 机制）

# Recover捕获异常
如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil。

```go
func Parse(input string) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)
        }
    }
}
```


