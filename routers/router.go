package routers

import (
	"github.com/astaxie/beego"
	"ls-tank-server-go/controllers"
)

func init() {
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Add")
	beego.Router("/user", &controllers.UserController{}, "get:Get")
	beego.Router("/user", &controllers.UserController{}, "patch:Update")
}
