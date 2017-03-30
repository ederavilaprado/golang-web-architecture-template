package daos

import (
	"fmt"

	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/ederavilaprado/golang-web-architecture-template/models"
	"github.com/jmoiron/sqlx"
)

type CustomerDAO struct {
	db *sqlx.DB
}

func NewCustomerDAO(db *sqlx.DB) *CustomerDAO {
	return &CustomerDAO{db}
}

func (dao *CustomerDAO) Get(rs app.RequestScope, ID int) (*models.Customer, error) {
	customer := &models.Customer{}

	fmt.Printf("DAO => %+v\n", rs.RequestID())

	if err := dao.db.Get(customer, "SELECT * FROM customers WHERE id=$1", ID); err != nil {
		return nil, err
	}
	return customer, nil
}
