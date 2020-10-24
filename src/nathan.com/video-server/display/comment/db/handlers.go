package db

import "stream-player/src/nathan.com/video-server/display/consts"

func AddNewComment(vid int, aid int, content string) error {
	stmtIns, err := util.dbConn.Prepare("insert into comments (video_id, author_id, content) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmtIns.Exec(vid, aid, content); err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func listComments(vid int, from, to int) ([]*consts.CommentInfo, error) {
	stmtOut, err := util.dbConn.Prepare(`select comments.id, users.username, comments.content
		from comments inner join users on comments.author_id = users.id 
		where comments.video_id = ? and comments.create_time in (from_unixtime(?), from_unixtime(?))`)
	var res []*consts.CommentInfo
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var vid int
		var name, content string
		if err := rows.Scan(&vid, &name, &content); err != nil {
			return res, err
		}
		c := &consts.CommentInfo{
			VideoId:    vid,
			AuthorName: name,
			Content:    content,
		}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
