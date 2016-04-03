package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "ls-tank-server-go/routers"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://127.0.0.1:*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Cache-Control", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	beego.Run()
}
