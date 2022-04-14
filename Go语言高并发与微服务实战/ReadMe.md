[TOC]

# 内容
- 语言基础
- 进阶
- Go web
- 并发、锁
    - 并发模型
    - CSP
    - 常见线程模型
    - MPG线程模型
- 微服务
    - 云原生与微服务
    - 分布式配置中心
        - Raft
        - Zookeeper
    - 服务注册与发现
    - 通信机制
    - 负载均衡
    - 容错处理
    - 分布式链路追踪
- 微服务框架: Go-kit、Go Micro、Java Spring Cloud、Node Seneca


# 微服务
- 将明确定义的功能分成更小的服务，并让每个服务独立迭代
- 架构：
    - 侵入式架构: 服务框架嵌入程序代码，组合各种组件，如RPC、LB
    - 非侵入式: 代理的形式

# SOA
- 面向服务架构: 模块化开发、分布式扩展部署、服务接口定义

# 常见微服务框架
- Java:
    - Spring Cloud
    - Dubbo
- Go:
    - Go-kit
    - Go Micro

# go mod
- go mod download: 下载依赖

# 并发模型
- CSP
- 常见线程模型
- MPG线程模型