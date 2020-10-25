package worker

import (
	"log"
	"time"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			log.Printf("clear job will start:%v", time.Now())
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(3, true, VideoClearDispatch, VideoClearExecutor)
	worker := NewWorker(10, r)
	worker.startWorker()
}
