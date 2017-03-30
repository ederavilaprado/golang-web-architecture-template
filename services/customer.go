package services

import (
	"fmt"

	"github.com/ederavilaprado/golang-web-architecture-template/models"
)

type customerDAO interface {
	// TODO: change to return model, not string...
	Get(ID int) (*models.Customer, error)
}

type CustomerService struct {
	dao customerDAO
}

func NewCustomerService(dao customerDAO) *CustomerService {
	return &CustomerService{dao}
}

func (s *CustomerService) Get(ID int) (*models.Customer, error) {

	fmt.Printf("=> %+v\n", "Inside CustomerService.Get...")
	return s.dao.Get(ID)

}
