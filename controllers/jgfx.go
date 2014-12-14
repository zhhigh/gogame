package controllers

import (
	"github.com/astaxie/beego"
)

type Jgfx struct {
	beego.Controller
}

func (c *Jgfx) Get() {
	c.TplNames = "jgfx.html"
}
