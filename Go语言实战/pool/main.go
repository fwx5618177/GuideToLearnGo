package main

import (
	"log"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"pool/pool"
)

const (
	maxCoroutines = 25
	pooledResoureces = 2 // 池中的资源的数量
)

// 每个连接分配一个独一无二的id
var idCounter int32

// dbConnection模拟共享的资源
type dbConnection struct {
	ID int32
}

// 完成对资源的释放
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)

	return nil
}

// 创建新的连接
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxCoroutines)

	p, err := pool.New(createConnection, pooledResoureces)

	if err != nil {
		log.Println(err)
	}

	// 池连接查询
	for query := 0; query < maxCoroutines; query++ {
		// 每个goroutine都要复制一份要查询的副本，不然所有的查询都会共享同一个查询变量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// 等待goroutine结束
	wg.Wait()

	// 关闭池
	log.Println("Shutdown Program.")

	p.Close()
}

// 测试链接的资源池
func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()

	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	// 等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	log.Println("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
