// 通过通道来监视程序的执行时间
package runner

// runner管理处理任务的运行和生命周期

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// 在给定的超时时间内执行一组任务,并在操作系统发送中断信号时结束这些任务
type Runner struct {
	// 从os发送的信号
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	// 一组索引，顺序执行的函数
	tasks []func(int)
}

var ErrTimeout = errors.New("timeout...")

var ErrInterrupt = errors.New("ErrInterrupt...")

// 返回一个新的Runners
func New(d time.Duration) *Runner {
	return &Runner {
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

// add to runner, param is callback(id int)
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// start,并监视通道事件
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <- r.complete:
		return err
	case <- r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测os的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}


func (r *Runner) gotInterrupt() bool {
	select {
	case <- r.interrupt:
		signal.Stop(r.interrupt)

		return true
	default:
		return false
	}


}