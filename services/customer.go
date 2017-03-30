package services

import (
	"fmt"

	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/models"
)

type customerDAO interface {
	Get(rs app.RequestScope, ID int) (*models.Customer, error)
}

type CustomerService struct {
	dao customerDAO
}

func NewCustomerService(dao customerDAO) *CustomerService {
	return &CustomerService{dao}
}

func (s *CustomerService) Get(rs app.RequestScope, ID int) (*models.Customer, error) {

	fmt.Printf("=> %+v\n", "Inside CustomerService.Get...")
	return s.dao.Get(rs, ID)

}
