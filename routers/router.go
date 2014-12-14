package routers

import (
	"gotmp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/jgfx", &controllers.MainController{})//激光防线
}
