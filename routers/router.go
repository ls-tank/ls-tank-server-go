package routers

import (
	"github.com/astaxie/beego"
	"ls-tank-server-go/controllers"
)

func init() {
	beego.Router("/api/user", &controllers.UserController{}, "post:Add")
	beego.Router("/api/user", &controllers.UserController{}, "get:Get")
	beego.Router("/api/user", &controllers.UserController{}, "patch:Update")
}
