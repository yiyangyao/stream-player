package taskrunner

import (
	"errors"
	"log"
	"os"
	"stream-player/src/nathan.com/video-server/scheduler/db"
	"sync"
)

func deleteVideoFile(videoName string) error {
	if err := os.Remove("./videos/" + videoName); err != nil {
		log.Printf("delete video file error: %v", err)
		return err
	}

	return nil
}

func VideoClearDispatch(dc dataChan) error {
	res, err := db.ReadVideoDeletionRecord(10)
	if err != nil {
		log.Printf("video clear dispatcher err: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("all tasks finished")
	}

	for _, name := range res {
		dc <- name
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case videoName := <-dc:
			go func(videoName interface{}) {
				if err := deleteVideoFile(videoName.(string)); err != nil {
					errMap.Store(videoName, err)
					return
				}
				if err := db.DelVideoDeletionRecord(videoName.(string)); err != nil {
					errMap.Store(videoName, err)
					return
				}
			}(videoName)
		default:
			break forloop
		}
	}
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
