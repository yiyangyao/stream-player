package db

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate videos")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestAddUserCredential)
	t.Run("Get", TestGetUser)
	t.Run("Del", TestDeleteUser)
	t.Run("ReGet", TestRegetUser)
}

func TestAddUserCredential(t *testing.T) {
	err := AddUserCredential("nathan", "nathan2012")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	passwd, err := getUserCredential("nathan")
	if passwd != "nathan2012" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func TestDeleteUser(t *testing.T) {
	err := deleteUser("nathan", "nathan2012")
	if err != nil {
		t.Errorf("Error of DelUser: %v", err)
	}
}

func TestRegetUser(t *testing.T) {
	passwd, err := getUserCredential("nathan")
	if err != nil {
		t.Errorf("Error of ReGEtUser: %v", err)
	}

	if passwd != "" {
		t.Errorf("Error of ReGEtUser")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	t.Run("PrepareUser", TestAddUserCredential)
	t.Run("Add", TestAddNewVideo)
	t.Run("Get", TestGetVideo)
	t.Run("Del", TestDeleteVideo)
	t.Run("ReGet", TestRegetVideo)
}

func TestAddNewVideo(t *testing.T) {
	video, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
	if video.AuthorId != 1 {
		t.Errorf("Error of AddUser: %v", video.AuthorId)
	}

}

func TestGetVideo(t *testing.T) {
	video, err := GetVideo("my-video")
	if video == nil || video.AuthorId != 1 || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func TestDeleteVideo(t *testing.T) {
	err := DeleteVideo("my-video")
	if err != nil {
		t.Errorf("Error of DelUser: %v", err)
	}
}

func TestRegetVideo(t *testing.T) {
	video, err := GetVideo("my-video")
	if err != nil {
		t.Errorf("Error of ReGetUser: %v", err)
	}

	if video.AuthorId != 0 {
		t.Errorf("Error of ReGEtUser: %v", video)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("adduser", TestAddUserCredential)
	t.Run("addcomments", TestAddNewComment)
	t.Run("listcomments", TestListComments)
}

func TestAddNewComment(t *testing.T) {
	vid := 12121212
	aid := 1
	content := "i like this video"
	if err := AddNewComment(vid, aid, content); err != nil {
		t.Errorf("%v", err)
	}
}

func TestListComments(t *testing.T) {
	vid := 12121212
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := listComments(vid, from, to)
	if err != nil {
		t.Errorf("%v", err)
	}
	for i, element := range res {
		fmt.Printf("conment: %d, %v \n", i, element)
	}
}
