package controllers

import (
	"github.com/astaxie/beego"
)

type jgfx struct {
	beego.Controller
}

func (c *jgfx) Get() {
	c.TplNames = "index.html"
}
