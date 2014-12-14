package controllers

import (
	"github.com/astaxie/beego"
)

type jgfx struct {
	beego.Controller
}

func (c *JgfxController) Get() {
	c.TplNames = "index.html"
}
