package controllers

import (
	"github.com/astaxie/beego"
)

type Jgfx struct {
	beego.Controller
}

func (c *Jgfx) Get() {
	beego.Debug("this is trace for get")
	c.TplNames = "jgfx/index.html"
	beego.Debug("this is trace for post")

}
