package consts

type CommentInfo struct {
	CommentId   int    `json:"comment_id"`
	VideoId     int    `json:"video_id"`
	AuthorId    int    `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Content     string `json:"content"`
	CommentTime string `json:"comment_time"`
}
