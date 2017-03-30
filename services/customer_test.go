package services

import "testing"
import "github.com/stretchr/testify/assert"

func TestCustomerServiceGet(t *testing.T) {
	s := NewCustomerService()

	customer, err := s.Get(123)

	if assert.Nil(t, err) && assert.NotNil(t, customer) {
		assert.Equal(t, "Eder Ávila Prado", customer.Name)
	}
}
