package defs

// requests
type UserCredential struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// Data Video
type VideoInfo struct {
	AuthorId  int
	VideoName string
}

// Data Comment
type CommentInfo struct {
	VideoId    int
	AuthorName string
	Content    string
}
