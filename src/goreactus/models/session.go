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
	UserId     	string
}


func AddSession(s Session) string {
	s.Id = "session_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	SessionList[s.Id] = &s
	return s.Id
}

func CreateSession(userid string) string {
	var session Session
	session.UserId = userid
	return AddSession(session)
}

func GetSession(sid string) (s *Session, err error) {
	if s, ok := SessionList[sid]; ok {
		return s, nil
	}
	return nil, errors.New("Session does not exist")
}

func GetAllSessions() map[string]*Session {
	return SessionList
}


func DeleteSession(sid string) {
	delete(SessionList, sid)
}
