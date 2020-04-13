package gothpool

type ExecPool struct {
	limiter chan bool
	queue   chan func()
	on      bool
}

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
			if !ep.on {
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

func (ep *ExecPool) Run(f func()) {
	ep.limiter <- true
	ep.queue <- f
}
