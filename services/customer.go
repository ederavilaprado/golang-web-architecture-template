package services

import "fmt"

type customerDAO interface {
	// TODO: change to return model, not string...
	Get(id int) (string, error)
}

type CustomerService struct {
	// TODO: DAO must came here
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) Get(id int) (string, error) {
	return fmt.Sprintf("Inside CustomerService.Get ID: %d", id), nil
}
