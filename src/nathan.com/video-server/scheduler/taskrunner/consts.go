package taskrunner

type controlChan chan string    // control chan
type dataChan chan interface{}  // data chan: goroutine生产和消费的数据管道
type fn func(dc dataChan) error // 任务执行函数

const (
	READY_TO_DISPATCH string = "d"
	READY_TO_EXECUTE  string = "e"
	CLOSE             string = "c"
)

const LOAD_VIDEO_RECORD_BUFFER int = 10
