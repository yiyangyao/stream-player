package taskrunner

type Runner struct {
	Controller controlChan // 控制信息管道(chan string): make(chan string, 1)
	Error      controlChan // 错误信息管道(chan string): make(chan string, 1)
	Data       dataChan    // 数据管道(chan interface{}): make(chan interface{}, dataSize)
	dataSize   int         // 数据管道的缓存大小
	longLived  bool        // 是否常驻
	Dispatcher fn          // 生产者(func(dc dataChan) error)
	Executor   fn          // 消费者(func(dc dataChan) error)
}

func NewRunner(size int, longLived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		dataSize:   size,
		longLived:  longLived,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				if err := r.Dispatcher(r.Data); err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				if err := r.Executor(r.Data); err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
