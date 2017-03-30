package services

import (
	"errors"
	"testing"

	"github.com/ederavilaprado/golang-web-architecture-template/models"
	"github.com/stretchr/testify/assert"
)

func TestCustomerServiceGet(t *testing.T) {
	m := newCustomerDAOMock()
	s := NewCustomerService(m)

	customer, err := s.Get(1)
	if assert.Nil(t, err) && assert.NotNil(t, customer) {
		assert.Equal(t, "Frances Jordan", customer.Name)
	}

	customer, err = s.Get(999)
	assert.NotNil(t, err, "999 should not be an valid customer ")
	assert.Nil(t, customer)
}

type customerDAOMock struct {
	records []*models.Customer
}

func newCustomerDAOMock() customerDAO {
	m := &customerDAOMock{}
	// Do you want to know from where this names come from...
	// https://www.mockaroo.com/ ;)
	m.records = []*models.Customer{
		&models.Customer{ID: 1, Name: "Frances Jordan"},
		&models.Customer{ID: 2, Name: "Ruth Woods"},
		&models.Customer{ID: 3, Name: "Doris Lynch"},
		&models.Customer{ID: 4, Name: "Annie Little"},
		&models.Customer{ID: 5, Name: "Dennis Richards"},
	}
	return m
}

func (d *customerDAOMock) Get(ID int) (*models.Customer, error) {
	for _, record := range d.records {
		if record.ID == ID {
			return record, nil
		}
	}
	return nil, errors.New("Customer not found")
}
