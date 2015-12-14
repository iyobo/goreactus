package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["goreactus/controllers:ActivityController"] = append(beego.GlobalControllerRouter["goreactus/controllers:ActivityController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:ActivityController"] = append(beego.GlobalControllerRouter["goreactus/controllers:ActivityController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:ActivityController"] = append(beego.GlobalControllerRouter["goreactus/controllers:ActivityController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goreactus/controllers:UserController"] = append(beego.GlobalControllerRouter["goreactus/controllers:UserController"],
		beego.ControllerComments{
			"Unauthorized",
			`/unauthorized`,
			[]string{"get"},
			nil})

}
