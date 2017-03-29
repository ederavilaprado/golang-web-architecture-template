package routes

import iris "gopkg.in/kataras/iris.v6"

// MapRoutes map all routes for the app
func MapRoutes(app *iris.Framework) {

	r := app.Party("/customers")

	mapCustomerRoutes(r)

}
