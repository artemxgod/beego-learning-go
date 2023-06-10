package routers

import (
	"learningApp/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:SayHello")
}
