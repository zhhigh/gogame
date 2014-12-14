package routers

import (
	"gogame/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/jgfx", &controllers.JgfxController{})//激光防线q
}
