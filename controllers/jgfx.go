package controllers

import (
	"github.com/astaxie/beego"
)

type Jgfx struct {
	beego.Controller
}

func (c *Jgfx) Get() {
	beego.Trace("this is trace for get")
	c.TplNames = "jgfx.html"

}
