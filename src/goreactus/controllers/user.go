package controllers

import (
	"goreactus/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type Response struct {
	success bool
	msg		string

}
// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"success":"true","uid": uid}
	u.ServeJson()
}

// @Title Get
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJson()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"success":"false","msg":err.Error()}
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJson()
}

// @Title update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		_, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = map[string]string{"success":"false","msg":err.Error()}
		} else {
			u.Data["json"] = map[string]string{"success":"true"}
		}
	}
	u.ServeJson()
}

// @Title delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = map[string]string{"success":"true","msg":"delete success!"}
	u.ServeJson()
}

// @Title login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	success, token, userid := models.Login(username, password)
	if success==true {
		u.Data["json"] = map[string]string{"success":"true","msg":"Success", "token":token,"username":username,"userid":userid}
	} else {
		u.Data["json"] = map[string]string{"success":"false","msg":"Invalid Credentials"}
	}
	u.ServeJson()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	sid := u.GetString("sid")
	models.DeleteSession(sid)

	u.Data["json"] = map[string]string{"success":"true","msg":"success"}
	u.ServeJson()
}

// @Title unauthorized
// @Description Logs out current logged in user session
// @Success 403 {string} unauthorized
// @router /unauthorized [get]
func (u *UserController) Unauthorized() {

	u.Data["json"] = map[string]string{"msg":"Unauthorized"}
	u.ServeJson()
}