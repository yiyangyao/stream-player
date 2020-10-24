package worker

import (
	"errors"
	"log"
	"os"
	"stream-player/src/nathan.com/video-server/scheduler/db"
	"sync"
)

func deleteVideoFile(videoName string) error {
	if err := os.Remove(VIDEO_DIR + videoName); err != nil {
		log.Printf("delete video file error: %v", err)
		return err
	}

	return nil
}

func VideoClearDispatch(dc dataChan) error {
	videoNames, err := db.ReadVideoDeletionRecord(LOAD_VIDEO_RECORD_BUFFER)
	if err != nil {
		log.Printf("read video deletion record err: %v", err)
		return err
	}

	if len(videoNames) == 0 {
		return errors.New("all tasks finished")
	}

	for _, name := range videoNames {
		dc <- name
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forLoop:
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
			break forLoop
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
