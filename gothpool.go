package gothpool

import "errors"

type ExecPool struct {
	limiter chan bool
	queue   chan func()
	on      bool
}

var ExecPoolStoppedErr = errors.New("executor pool is stopped")

func New(parallelism int64, queueSize int64) *ExecPool {
	return &ExecPool{
		limiter: make(chan bool, parallelism),
		queue:   make(chan func(), queueSize),
	}
}

func (ep *ExecPool) Start() {
	ep.on = true
	go func() {
		for {
			if !ep.on && len(ep.queue) == 0 {
				break
			}
			select {
			case job := <-ep.queue:
				job()
				<-ep.limiter
			}
		}
	}()
}

func (ep *ExecPool) Stop() {
	ep.on = false
}

func (ep *ExecPool) Run(f func()) error {
	if !ep.on {
		return ExecPoolStoppedErr
	}
	ep.limiter <- true
	ep.queue <- f
	return nil
}
