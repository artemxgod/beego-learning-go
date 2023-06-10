package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

// can define your own controller
type TestController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// outputs hello world
func (hello *MainController) SayHello() {
	// pass variable from controller to template
	hello.Data["name"] = "Alex"
	// setting template name for receiver
	hello.TplName = "hello.html"
	// we then create a template in /view with all the contect
	// and now we can use this template in router
}