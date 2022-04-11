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



