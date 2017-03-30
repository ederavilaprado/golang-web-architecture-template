package main

import (
	"log"

	"github.com/ederavilaprado/golang-web-architecture-template/apis"
	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/daos"
	"github.com/ederavilaprado/golang-web-architecture-template/services"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	irisApp.Use(app.Init())

	// Creates a router group
	rg := irisApp.Party("/v0")

	// Starting DB...
	db, err := sqlx.Connect("postgres", "user=postgres password=mysecretpassword dbname=apidb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	// TODO: put defer db.Close() somewhere
	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(10)

	customerDAO := daos.NewCustomerDAO(db)
	apis.ServeCustomerResource(rg, services.NewCustomerService(customerDAO))

	return irisApp
}
