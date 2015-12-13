package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	SessionList map[string]*Session
)

func init() {
	SessionList = make(map[string]*Session)
}

//A typical full session object would have expiry, ip, etc. But this is a volatile DB so...
type Session struct {
	Id       	string
	User     	User
}


func AddSession(s Session) string {
	s.Id = "session_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	SessionList[s.Id] = &s
	return s.Id
}

func CreateSession(u User) string {
	var session Session
	session.User = u
	return AddSession(session)
}

func GetSession(uid string) (s *Session, err error) {
	if s, ok := SessionList[uid]; ok {
		return s, nil
	}
	return nil, errors.New("Session does not exist")
}

func GetAllSessions() map[string]*Session {
	return SessionList
}


func DeleteSession(uid string) {
	delete(SessionList, uid)
}
