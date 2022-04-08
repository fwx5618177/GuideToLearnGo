package pool
// 管理用户定义的一组资源

import (
	"errors"
	"log"
	"io"
	"sync"
)

// pool管理一组可以安全在多个goroutine间共享的资源
// 被管理的资源必须实现io.Close接口
type Pool struct {
	m sync.Mutex
	resources chan io.Closer
	factory func() (io.Closer, error)
	closed bool
}

// ErrPoolClosed-Acquire one closed pool
var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value is empty.")
	}

	return &Pool{
		factory: fn,
		resources: make(chan io.Closer, size),
	}, nil
}


// acquire
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")

		if !ok {
			return nil, ErrPoolClosed
		}

		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// release
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()

	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
		// 资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	// 队列满了，关闭资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}


// close, shut down all
func (p *Pool) Close() {
	p.m.Lock()

	defer p.m.Unlock()

	if p.closed {
		return
	}

	// 关闭池
	p.closed = true

	// 清空前关闭通道,否则会发生死锁
	close(p.resources)

	// 关闭
	for r := range p.resources {
		r.Close()
	}
}