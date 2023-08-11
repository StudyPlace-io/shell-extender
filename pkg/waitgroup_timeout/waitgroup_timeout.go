package waitgroup_timeout

import (
	"sync"
	"time"
)

// WaitGroupWithTimeout 增加超时功能的WaitGroup
type WaitGroupWithTimeout struct {
	sync.WaitGroup
	timeout time.Duration // 超时时间
}

func NewWaitGroupWithTimeout(timeout time.Duration) *WaitGroupWithTimeout {
	w := &WaitGroupWithTimeout{
		sync.WaitGroup{},
		timeout,
	}
	return w
}

// WaitTimeout 判断Wait() 是否超时
func (wg *WaitGroupWithTimeout) WaitTimeout() bool {

	ch := make(chan bool, 1)

	go time.AfterFunc(wg.timeout, func() {
		ch <- true
	})

	go func() {
		wg.Wait()
		ch <- false
	}()

	return <-ch
}
