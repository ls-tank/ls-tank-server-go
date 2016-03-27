package routers

import (
	"github.com/astaxie/beego"
	"ls-tank-server-go/controllers"
)

func init() {
	beego.Router("/api/login", &controllers.UserController{}, "post:Login")
	beego.Router("/api/register", &controllers.UserController{}, "post:Add")
	beego.Router("/api/user", &controllers.UserController{}, "get:Get")
	beego.Router("/api/user", &controllers.UserController{}, "patch:Update")
}
