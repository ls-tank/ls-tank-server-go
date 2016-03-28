package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "ls-tank-server-go/routers"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Cache-Control", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	beego.Run()
}
