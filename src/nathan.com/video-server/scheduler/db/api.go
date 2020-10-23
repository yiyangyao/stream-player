package db

import "database/sql"

func AddVideoDeletionRecord(vid int) error {
	stmtOut, err := dbConn.Prepare("select video_name from videos where id = ?")
	if err != nil {
		return err
	}
	var videoName string
	if err = stmtOut.QueryRow(vid).Scan(&videoName); err != nil && err != sql.ErrNoRows {
		return err
	}
	defer stmtOut.Close()
	stmtIns, err := dbConn.Prepare("INSERT INTO video_to_be_deleted (video_id, video_name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	if _, err = stmtIns.Exec(vid, videoName); err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
