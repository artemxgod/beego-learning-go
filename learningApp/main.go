package main

import (
	"learningApp/controllers"
	_ "learningApp/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
