package routers

import (
	"learningApp/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:SayHello")
	beego.Router("/home", &controllers.MainController{}, "get:GetHome")
			// adding a special handler for a not found error
	beego.ErrorController(&controllers.MainController{})
}

