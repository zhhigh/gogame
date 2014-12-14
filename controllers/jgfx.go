package controllers

import (
	"github.com/astaxie/beego"
)

type JgfxController struct {
	beego.Controller
}

func (c *JgfxController) Get() {
	c.TplNames = "jgfx.html"
}
