目录：
[TOC]

# Go并发编程
- 工程规范
- 工具链
- 并发编程模型和机制

书结构:
1. 优缺点
2. 工程结构
3. 测试和文档要求
4. 并发模型

# 优缺点
缺点：
1. 分布式计算不如Erlang
2. 速度不如C
3. 垃圾回收采用的是标记清除法，在回收期间停止所有用户程序的操作

# 工程结构
- src
- pkg
- bin
- GOPATH: `export GOPATH=$HOME`
- go build: main下
- go install
- 代理: go env -w GOPROXY=https://goproxy.cn,direct

# 高级数据结构
- Set

# 程序测试和文档
- 添加文档
1. 测试函数: `func TestXxxs(t *testing.T)`
2. 常规记录: `t`的Log / Logf记录常规信息
    - Log = fmt.Println
    - Logf = fmt.Printf
3. 错误记录: `t`的Error / Errorf
4. 致命错误: `t` Fatal / Fatalf
5. 失败标记: t.Fail
6. 立即失败标记: t.FailNow
7. 失败判断: t.Failed
8. 忽略测试: t.SkipNow
9. 并行运行: t.Parallel
10. 功能测试的运行: go test

# 多进程编程
- ICP: 进程间通讯
- ICP三类：
    - 基于通讯
    - 基于信号: os的信号
    - 基于同步: 信号灯
- 基于通讯ICP:
    - 数据传送
        - Pipe
        - Message Queue
    - 共享内存



