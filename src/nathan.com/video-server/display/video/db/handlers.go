package db

import (
	"database/sql"
	"log"
	"stream-player/src/nathan.com/video-server/display/util"
	"stream-player/src/nathan.com/video-server/display/video/consts"
)

func AddNewVideo(authorId int, videoName string) (*consts.VideoInfo, error) {
	stmtIns, err := util.DBConn.Prepare("insert into videos (author_id, video_name) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(authorId, videoName)
	if err != nil {
		return nil, err
	}
	video := &consts.VideoInfo{
		AuthorId:  authorId,
		VideoName: videoName,
	}
	return video, nil
}

func GetVideoInfo(name string) (*consts.VideoInfo, error) {
	stmtOut, err := util.DBConn.Prepare("select video_id, author_id from videos where video_name = ?")
	if err != nil {
		return nil, err
	}
	var videoId int
	var authorId int
	err = stmtOut.QueryRow(name).Scan(&videoId, &authorId)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()
	video := &consts.VideoInfo{
		VideoId:   videoId,
		AuthorId:  authorId,
		VideoName: name,
	}
	return video, nil
}

func DeleteVideo(name string) error {
	stmtDel, err := util.DBConn.Prepare("delete from videos where video_name = ?")
	if err != nil {
		log.Printf("%s", err)
	}
	_, err = stmtDel.Exec(name)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}

func ListVideoInfo(authorId int) ([]*consts.VideoInfo, error) {
	stmtOut, err := util.DBConn.Prepare("select video_id, video_name from videos where author_id = ?")
	if err != nil {
		return nil, err
	}

	var videoList []*consts.VideoInfo
	rows, err := stmtOut.Query(authorId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var videoId int
		var videoName string
		if err := rows.Scan(&videoId, &videoName); err != nil {
			return nil, err
		}
		video := &consts.VideoInfo{
			VideoId:   videoId,
			AuthorId:  authorId,
			VideoName: videoName,
		}
		videoList = append(videoList, video)
	}
	defer stmtOut.Close()

	return videoList, nil

}
