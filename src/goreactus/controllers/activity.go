package controllers

import (
	"goreactus/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Activity
type ActivityController struct {
	beego.Controller
}


// @Title createActivity
// @Description create activity
// @Param	body		body 	models.Activity	true		"body for activity content"
// @Success 200 {string} models.Activity.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ActivityController) Post() {
	var activity models.Activity
	json.Unmarshal(c.Ctx.Input.RequestBody, &activity)
	uid := models.AddActivity(activity)
	c.Data["json"] = map[string]string{"uid": uid}
	c.ServeJson()
}

// @Title Get
// @Description get all Activitys
// @Success 200 {object} models.Activity
// @router / [get]
func (c *ActivityController) GetAll() {
	activitys := models.GetAllActivity()
	c.Data["json"] = activitys
	c.ServeJson()
}

// @Title Get
// @Description get activity by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Activity
// @Failure 403 :uid is empty
// @router /:uid [get]
func (c *ActivityController) Get() {
	uid := c.GetString(":uid")
	if uid != "" {
		activity, err := models.GetActivity(uid)
		if err != nil {
			c.Data["json"] = err
		} else {
			c.Data["json"] = activity
		}
	}
	c.ServeJson()
}

// @Title delete
// @Description delete the activity
// @Param	uid		path 	string	true		"The activity id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (c *ActivityController) Delete() {
	uid := c.GetString(":uid")
	models.DeleteActivity(uid)
	c.Data["json"] = "delete success!"
	c.ServeJson()
}
