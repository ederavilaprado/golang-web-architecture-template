package routes

import iris "gopkg.in/kataras/iris.v6"

func mapCustomerRoutes(r *iris.Router) {

	r.Get("", func(ctx *iris.Context) {
		ctx.Text(200, "Just a text here")
	})

}

// func CustomerRoutes() *iris.Router {

// 	getOneHandler := func(ctx *iris.Context) {
// 		ctx.Text(200, "Just a text here")
// 	}
//
// 	// r := *iris.Router{}
//
// 	// r.Get("/:id", getOneHandler)
//
//
//
// 	return r
//
// }
