package consts

// request
type UserCredential struct {
	UserID   int    `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type SimpleSession struct {
	UserName string
	TTL      int64
}
