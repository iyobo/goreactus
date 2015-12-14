package controllers

import (
	"goreactus/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

// Operations about Activity
type ActivityController struct {
	beego.Controller
}

type ActivityBody struct {
	Name	string
	Change	string
}

// @Title createActivity
// @Description create activity
// @Param	body		body 	models.Activity	true		"body for activity content"
// @Success 200 {string} models.Activity.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ActivityController) Post() {
	var body ActivityBody
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	models.LogActivity(body.Name)
	c.Data["json"] = "ok"
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

// @Title delete
// @Description delete the activity
// @Param	uid		path 	string	true		"The activity id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (c *ActivityController) Delete() {
	name := c.GetString(":name")
	models.DeleteActivity(name)
	c.Data["json"] = "delete success!"
	c.ServeJson()
}
