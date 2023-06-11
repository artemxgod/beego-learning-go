# beego-learning-go

- In this project i will learn how to use beego framework in golang project

## What is beego

- Beego is a RESTful HTTP framework that can be used to build
    * APIs
    * web-apps
    * backend services
- It is based on MVC architecture

### MVC architecture

- `M.V.C.` stands for model, view and controller
- `Model` corresponds to all the data related logic
- `View` is used for UI logic of the app
- `Controller` contains the business logic

### Simple application using beego

```go
package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
```

- Without any configuration server will be initiated on localhost:8080
- And we will receive a following message in console: 'init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: The system cannot find the path specified'

### Installing bee tool

    go install github.com/beego/bee/v2@latest

- This tool allows us to create architecture for our app. Simply run:

```cmd
    $ bee new <application_name>
```

- run go mod tidy and then run the app
```cmd
    $ go mod tidy
    $ bee run
```
- visit `localhost:8080` to check the app


## Layout

- `/conf` - contains basic configuration of the application
- `/controllers` - can specify your own controllers. controllers manage backend pages and implement features
- `/models` - contains data models for db
- `/routers` - allows to create routes for web application
- `/static` - keeps static files like .css .js and images
- `/tests` - testing application methods
- `/views` template files

- [!] beego restarts server after modification

### template parsing 

- To pass variable from controller to template
    * Set variable in controller
    * parse it in the template
- Example:

```go
func (hello *MainController) SayHello() {
// setting template name
hello.TplName = "hello.html"
//setting a variable to controller.Data map
hello.Data["name"] = "Alex"
}
```

```html
<!-- hello.html -->
<h2>Hello Golang from {{.name}}</h2>
```

- Change syntax for variable using:
```go
// main.go
beego.BConfig.WebConfig.TemplateLeft = "<<<"
beego.BConfig.WebConfig.TemplateRight = ">>>"
```
- `<<<>>>` will replace `{{}}` syntax


- Change `/views` path to custom:
```go
// main.go
beego.AddViewPath("path_name")

// and add path in controller's method

controller.ViewPath = "path_name"
```

### Errors

- You can define a special handler for an specified error code using:
    * `ErrorHandler(code, handler)` if you define handler
    * `ErrorController(controller)` if you define controller method

- [!] Controller's method name should match form: "Error + error code". Example: `Error404()` otherwise it won't be recognised properly
