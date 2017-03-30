package services

import (
	"fmt"

	"github.com/ederavilaprado/golang-web-architecture-template/models"
)

type customerDAO interface {
	// TODO: change to return model, not string...
	Get(ID int) (models.Customer, error)
}

type CustomerService struct {
	// TODO: DAO must came here
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) Get(ID int) (*models.Customer, error) {

	c := &models.Customer{}
	c.ID = ID
	c.Name = "Eder Ávila Prado"

	fmt.Printf("=> %+v\n", "Inside CustomerService.Get...")

	return c, nil
}
