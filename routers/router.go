package routers

import (
	"github.com/astaxie/beego"
	"ls-tank-server-go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
