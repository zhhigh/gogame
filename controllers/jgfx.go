package controllers

import (
	"github.com/astaxie/beego"
)

type Jgfx struct {
	beego.Controller
}

func (c *jgfx) Get() {
	c.TplNames = "index.html"
}
