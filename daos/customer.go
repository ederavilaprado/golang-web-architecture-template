package daos

import (
	"github.com/ederavilaprado/golang-web-architecture-template/models"
	"github.com/jmoiron/sqlx"
)

type CustomerDAO struct {
	db *sqlx.DB
}

func NewCustomerDAO(db *sqlx.DB) *CustomerDAO {
	return &CustomerDAO{db}
}

func (dao *CustomerDAO) Get(ID int) (*models.Customer, error) {
	customer := &models.Customer{}
	if err := dao.db.Get(customer, "SELECT * FROM customers WHERE id=$1", ID); err != nil {
		return nil, err
	}
	return customer, nil
}
