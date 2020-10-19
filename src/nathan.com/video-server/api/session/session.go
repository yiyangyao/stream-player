package session

import (
	"log"
	"stream-player/src/nathan.com/video-server/api/db"
	"stream-player/src/nathan.com/video-server/api/defs"
	"sync"
	"time"

	"github.com/google/uuid"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	r, err := db.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})

}

func CreateNewSessionId(username string) string {
	// V1 基于时间
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	ct := time.Now().UnixNano() / 1000000
	ttl := ct + 30*60*1000 // 30 min
	ss := &defs.SimpleSession{UserName: username, TTL: ttl}
	db.InsertSession(uid.String(), ttl, username)
	sessionMap.Store(uid, ss)
	return uid.String()
}

func IsSessionExpired(sessionId string) (string, bool) {
	ss, ok := sessionMap.Load(sessionId)
	if ok {
		ct := time.Now().UnixNano() / 1000000
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sessionId)
			return "", true
		}

		return ss.(*defs.SimpleSession).UserName, false
	}

	return "", true
}

func deleteExpiredSession(sid string) {
	if err := db.DeleteSession(sid); err != nil {
		return
	}
	sessionMap.Delete(sid)
}
