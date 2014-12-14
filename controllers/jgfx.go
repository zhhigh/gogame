package controllers

import (
	"github.com/astaxie/beego"
)

type Jgfx struct {
	beego.Controller
}

func (c *Jgfx) Get() {
	beego.Debug("this is trace for get")
	c.TplNames = "index.html"
	//c.TplNames = c.getTplFileName("login")
	//this.TplNames = this.getTplFileName("login")
	beego.Debug("this is trace for post")

}
