package routers

import (
	"github.com/astaxie/beego"
	"ls-tank-server-go/controllers"
)

func init() {
	beego.Router("/api/user", &controllers.UserController{}, "post:Add")
	beego.Router("/api/user/:id", &controllers.UserController{}, "get:Get")
}
