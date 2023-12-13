package datastore 

import (
	"car-service/model"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type CustomerDatastore struct{}

func NewCustomerDatastore() *CustomerDatastore {
	return &CustomerDatastore{}
}

func (d *CustomerDatastore) Create(ctx *gofr.Context, customer *model.Customer) (*model.Customer, error) {
	var resp model.Customer

	queryInsert := "INSERT INTO customers (id, name) VALUES (?, ?)"

	result, err := ctx.DB().ExecContext(ctx, queryInsert, customer.ID, customer.Name)

	if err != nil {
		return &model.Customer{}, errors.DB{Err: err}
	}
	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return &model.Customer{}, errors.DB{Err: err}
	}

	querySelect := "SELECT id, name FROM customers WHERE id = ?"

	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).
		Scan(&resp.ID, &resp.Name)

	if err != nil {
		return &model.Customer{}, errors.DB{Err: err}
	}

	return &resp, nil
}


func (d *CustomerDatastore) GetAll(ctx *gofr.Context) ([]model.Customer,error){
	var customers []model.Customer

	rows,err := ctx.DB().QueryContext(ctx, "SELECT id, name FROM customers")
	if err != nil {
		return nil, errors.DB{Err:err}
	}

	for rows.Next(){
		var customer model.Customer
		if err := rows.Scan(&customer.ID, &customer.Name); err != nil{
			return nil, errors.DB{Err:err}
		}
		customers = append(customers, customer)
	}
	return customers, nil
}