package apis

import (
	"fmt"

	iris "gopkg.in/kataras/iris.v6"
)

type customerService interface {
	Get(id int) (string, error)
}

type customerResource struct {
	service customerService
}

func (r *customerResource) get(ctx *iris.Context) {
	// TODO: parse params, etc...

	customerID, _ := ctx.ParamInt("id")

	customer, err := r.service.Get(customerID)

	fmt.Printf("=> %+v\n", customer)

	if err != nil {
		ctx.Text(500, err.Error())
	}
	ctx.Text(200, customer)
}

func ServeCustomerResource(router *iris.Router, s customerService) {
	resource := &customerResource{s}

	r := router.Party("/customers")
	r.Get("/:id", resource.get)
}
