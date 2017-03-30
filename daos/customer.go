package daos

import "github.com/ederavilaprado/golang-web-architecture-template/models"

type CustomerDAO struct{}

func NewCustomerDAO() *CustomerDAO {
	return &CustomerDAO{}
}

func (dao *CustomerDAO) Get(ID int) (*models.Customer, error) {
	customer := &models.Customer{}

	customer.ID = ID
	customer.Name = "Fake name here"

	// TODO: select should come here

	return customer, nil
}
