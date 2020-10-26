package db

import (
	"stream-player/src/nathan.com/video-server/display/comment/consts"
	"stream-player/src/nathan.com/video-server/display/util"
)

func AddNewComment(vid int, aid int, content string) error {
	stmtIns, err := util.DBConn.Prepare("insert into comments (video_id, author_id, content) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmtIns.Exec(vid, aid, content); err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListVideoComments(vid, from, to int) ([]*consts.CommentInfo, error) {
	stmtOut, err := util.DBConn.Prepare(
		`select comments.id, comments.author_id, comments.content, comments.create_time, users.username
		from comments inner join users on comments.author_id = users.id 
		where 
		comments.video_id = ? and 
		comments.create_time >= from_unixtime(?) and 
		comments.create_time <= from_unixtime(?)`)
	var videoComments []*consts.CommentInfo
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var commentId, authorId int
		var createTime, content, userName string
		if err := rows.Scan(&commentId, &authorId, &createTime, &content, &userName); err != nil {
			return videoComments, err
		}
		c := &consts.CommentInfo{
			CommentId:   commentId,
			VideoId:     vid,
			AuthorId:    authorId,
			AuthorName:  userName,
			Content:     content,
			CommentTime: createTime,
		}
		videoComments = append(videoComments, c)
	}
	defer stmtOut.Close()
	return videoComments, nil
}
