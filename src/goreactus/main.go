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
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
