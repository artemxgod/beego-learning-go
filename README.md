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
