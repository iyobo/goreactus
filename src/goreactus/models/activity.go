package models

import (

)

var (
	ActivityMap map[string]int
)

func init() {
	ActivityMap = make(map[string]int)
	ActivityMap["Eating"] = 3
	ActivityMap["Coding"] = 6
	ActivityMap["Biking"] = 1
	ActivityMap["Swimming"] = 1
	ActivityMap["Yelling"] = 2
	ActivityMap["Plotting"] = 9
}

//Adds a new activity or increments an existing one
func LogActivity(a string) {
	var cnt = ActivityMap[a]

	if cnt==0{
		//Create new label....or rather just ignore
		//ActivityMap[a] = 1
	}else{
		//increment
		ActivityMap[a] = cnt+1
	}

}

func GetAllActivity() map[string]int {
	return ActivityMap
}


func DeleteActivity(activity string) {
	delete(ActivityMap, activity)
}
