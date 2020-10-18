package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"stream-player/src/nathan.com/video-server/api/defs"
)

func AddUserCredential(username string, passwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (username, passwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(username, passwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func getUserCredential(username string) (string, error) {
	stmtOut, err := dbConn.Prepare("select passwd from users where username = ?")
	if err != nil {
		log.Printf("%s", err)
	}
	var passwd string
	err = stmtOut.QueryRow(username).Scan(&passwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()

	return passwd, nil
}

func deleteUser(username string, passwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where username = ? and passwd = ?")
	if err != nil {
		log.Printf("%s", err)
	}
	_, err = stmtDel.Exec(username, passwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	stmtIns, err := dbConn.Prepare("insert into videos (author_id, video_name) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(aid, name)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{
		AuthorId:  aid,
		VideoName: name,
	}
	return res, nil
}

func GetVideo(name string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id, video_name from videos where video_name = ?")
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
	res := &defs.VideoInfo{
		AuthorId:  author_id,
		VideoName: video_name,
	}
	return res, nil
}

func DeleteVideo(name string) error {
	stmtDel, err := dbConn.Prepare("delete from videos where video_name = ?")
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

func AddNewComment(vid int, aid int, content string) error {
	stmtIns, err := dbConn.Prepare("insert into comments (video_id, author_id, content) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmtIns.Exec(vid, aid, content); err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func listComments(vid int, from, to int) ([]*defs.CommentInfo, error) {
	stmtOut, err := dbConn.Prepare(`select comments.id, users.username, comments.content
		from comments inner join users on comments.author_id = users.id 
		where comments.video_id = ? and comments.create_time in (from_unixtime(?), from_unixtime(?))`)
	var res []*defs.CommentInfo
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
		c := &defs.CommentInfo{
			VideoId:    vid,
			AuthorName: name,
			Content:    content,
		}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
