package defs

// requests
type UserCredential struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
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

type SimpleSession struct {
	UserName string
	TTL      int64
}
