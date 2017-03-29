package main

import (
	"github.com/ederavilaprado/golang-web-architecture-template/routes"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

var (
	app *iris.Framework
)

func init() {
	app = iris.New()
	app.Adapt(httprouter.New())

}

func main() {

	routes.MapRoutes(app)

	app.Listen("localhost:8080")

}
