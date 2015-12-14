package main

import (
	_ "goreactus/docs"
	_ "goreactus/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

var RedirectHttp = func(ctx *context.Context) {
	if !ctx.Input.IsSecure() {
		// no need for an additional '/' between domain and uri
		url := "https://" + ctx.Input.Domain()+":"+beego.AppConfig.String("httpsport") + ctx.Input.Uri()
		ctx.Redirect(302, url)
	}
}

func main() {
	beego.InsertFilter("/", beego.BeforeRouter, RedirectHttp)  // for http://mysite
	beego.InsertFilter("*", beego.BeforeRouter, RedirectHttp)  // for http://mysite/*
	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//Auth middleware
	var AuthFilter = func(ctx *context.Context) {
		/*TODO: Lazy & unsafe session check. We're not checking for session token validity.
		 Only if the user is sending us a session token.
		 This was never meant to be a full application, neither should this demo work be used in a full application.
		 It is simply an expose of how Go works with React.
		 - Iyobo Eki
		 */
		auth := ctx.Input.Header("Authorization")
		if auth == "" {
			ctx.Redirect(302, "/v1/user/unauthorized")
		}
	}

	// No auth middleware for user or activity controller to allow for easy viewing from REST API without auth header
	beego.InsertFilter("/nevermind", beego.BeforeRouter, AuthFilter)
	//beego.InsertFilter("/*/activity", beego.BeforeRouter, AuthFilter)
	//beego.InsertFilter("/*/user", beego.BeforeRouter, AuthFilter)




	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

	//Run REST API Server
	beego.Run()
}
