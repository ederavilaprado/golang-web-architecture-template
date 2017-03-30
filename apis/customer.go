package apis

import (
	"github.com/ederavilaprado/golang-web-architecture-template/models"

	iris "gopkg.in/kataras/iris.v6"
)

type customerService interface {
	Get(id int) (*models.Customer, error)
}

type customerResource struct {
	service customerService
}

func (r *customerResource) get(ctx *iris.Context) {
	// TODO: parse params, etc...
	// TODO: use some validation helper for POST methods

	customerID, _ := ctx.ParamInt("id")

	customer, err := r.service.Get(customerID)

	if err != nil {
		ctx.Text(500, err.Error())
	}
	ctx.JSON(200, customer)
}

func ServeCustomerResource(router *iris.Router, s customerService) {
	resource := &customerResource{s}

	r := router.Party("/customers")
	r.Get("/:id", resource.get)
}
