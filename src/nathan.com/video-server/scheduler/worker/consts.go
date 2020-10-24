package worker

type controlChan chan string    // control chan
type dataChan chan interface{}  // data chan: goroutine生产和消费的数据管道
type fn func(dc dataChan) error // 任务执行函数

// chan control
const (
	READY_TO_DISPATCH string = "d"
	READY_TO_EXECUTE  string = "e"
	CLOSE             string = "c"
)

// 每次读取数据的条数
const LOAD_VIDEO_RECORD_BUFFER int = 10

const VIDEO_DIR string = "/Users/bytedance/stream-player/src/nathan.com/video-server/stream-server/sources/videos/"
