package datastore 

import (
	"database/sql"
	"strconv"
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



func (d *CustomerDatastore) GetByID(ctx *gofr.Context, id string) (*model.Customer,error){
	var resp model.Customer

	querySelect := "SELECT id, name FROM customers WHERE id = ?"

	err := ctx.DB().QueryRowContext(ctx, querySelect, id).Scan(&resp.ID, &resp.Name)

	switch err {
	case sql.ErrNoRows:
		return &model.Customer{},errors.EntityNotFound{Entity: "customer", ID:id}
	case nil:
		return &resp, nil
	default:
		return &model.Customer{}, errors.DB{Err: err}
	}
}


func (d *CustomerDatastore) Update(ctx *gofr.Context, customer *model.Customer) (*model.Customer, error){
    _, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET name = ? WHERE id = ?", customer.Name, customer.ID)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return d.GetByID(ctx, strconv.Itoa(customer.ID))
}


func (d *CustomerDatastore) Delete(ctx *gofr.Context, id string) error {
	queryDelete := "DELETE FROM customers WHERE id = ?"

	_, err := ctx.DB().ExecContext(ctx, queryDelete, id)
	
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil

}