package db

import (
	"database/sql"
	"log"
)

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

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("select video_name from video_to_be_deleted limit ?")

	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord err: %v", err)
		return nil, err
	}

	var names []string
	for rows.Next() {
		var videoName string
		if err := rows.Scan(&videoName); err != nil {
			return nil, err
		}
		if len(videoName) != 0 {
			names = append(names, videoName)
		}
	}

	defer stmtOut.Close()
	return names, nil
}

func DelVideoDeletionRecord(videoName string) error {
	stmtDel, err := dbConn.Prepare("delete from video_to_be_deleted where video_name = ?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(videoName)
	if err != nil {
		log.Printf("delete video err %v", err)
		return err
	}

	defer stmtDel.Close()

	return nil
}
