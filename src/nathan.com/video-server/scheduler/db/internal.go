package db

import "log"

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("select video_name from video_to_be_deleted")

	var names []string

	if err != nil {
		return names, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord err: %v", err)
		return names, err
	}

	for rows.Next() {
		var videoName string
		if err := rows.Scan(&videoName); err != nil {
			return names, err
		}

		names = append(names, videoName)
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
