package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	ActivityList map[string]*Activity
)

func init() {
	ActivityList = make(map[string]*Activity)
}

type Activity struct {
	Id       string
	Username string
	action	 string
}


func AddActivity(a Activity) string {
	a.Id = "activity_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	ActivityList[a.Id] = &a
	return a.Id
}

func CreateActivity(u string, action string) string {
	var activity Activity
	activity.Username = u
	activity.action = action
	return AddActivity(activity)
}

func GetActivity(uid string) (s *Activity, err error) {
	if s, ok := ActivityList[uid]; ok {
		return s, nil
	}
	return nil, errors.New("Activity does not exist")
}

func GetAllActivity() map[string]*Activity {
	return ActivityList
}


func DeleteActivity(uid string) {
	delete(ActivityList, uid)
}
