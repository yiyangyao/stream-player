package db

import (
	"database/sql"
	"log"
	"stream-player/src/nathan.com/video-server/display/consts"
	"stream-player/src/nathan.com/video-server/display/util"
)

func AddNewVideo(aid int, name string) (*consts.VideoInfo, error) {
	stmtIns, err := util.dbConn.Prepare("insert into videos (author_id, video_name) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(aid, name)
	if err != nil {
		return nil, err
	}
	res := &consts.VideoInfo{
		AuthorId:  aid,
		VideoName: name,
	}
	return res, nil
}

func GetVideo(name string) (*consts.VideoInfo, error) {
	stmtOut, err := util.dbConn.Prepare("select author_id, video_name from videos where video_name = ?")
	if err != nil {
		return nil, err
	}
	var author_id int
	var video_name string
	err = stmtOut.QueryRow(name).Scan(&author_id, &video_name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()
	res := &consts.VideoInfo{
		AuthorId:  author_id,
		VideoName: video_name,
	}
	return res, nil
}

func DeleteVideo(name string) error {
	stmtDel, err := util.dbConn.Prepare("delete from videos where video_name = ?")
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
