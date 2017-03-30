package main

import (
	"github.com/ederavilaprado/golang-web-architecture-template/apis"
	"github.com/ederavilaprado/golang-web-architecture-template/daos"
	"github.com/ederavilaprado/golang-web-architecture-template/services"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {

	// TODO: load config here
	// TODO: load config constants here

	// TODO: create log
	// TODO: create connection pool

	// TODO: build routes here

	app := buildRouter()
	app.Listen("localhost:8080")

}

func buildRouter() *iris.Framework {
	// TODO: should accept log + db... ?!?

	irisApp := iris.New()
	irisApp.Adapt(iris.DevLogger())
	irisApp.Adapt(httprouter.New())

	// TODO: new Customer DAO here
	// TODO: bind Customer DAO to service

	appRouter := irisApp.Party("/v0")

	customerDAO := daos.NewCustomerDAO()
	apis.ServeCustomerResource(appRouter, services.NewCustomerService(customerDAO))

	return irisApp
}
