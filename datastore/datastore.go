package datastore

import (
	"car-service/model"
	"database/sql"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type CustomerDatastore struct{}

func NewCustomerDatastore() *CustomerDatastore {
	return &CustomerDatastore{}
}

func (d *CustomerDatastore) Create(ctx *gofr.Context, customer *model.Customer) (*model.Customer, error) {
	var resp model.Customer

	queryInsert := "INSERT INTO customers (id, name, email, phone, address, city, date_of_birth, is_active) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := ctx.DB().ExecContext(ctx, queryInsert, customer.ID, customer.Name, customer.Email, customer.Phone, customer.Address, customer.City, customer.DateOfBirth, customer.IsActive)

	if err != nil {
		return &model.Customer{}, errors.DB{Err: err}
	}
	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return &model.Customer{}, errors.DB{Err: err}
	}

	querySelect := "SELECT id, name, email, phone, address, city, date_of_birth, is_active FROM customers WHERE id = ?"

	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).
		Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Phone, &resp.Address, &resp.City, &resp.DateOfBirth, &resp.IsActive)

	if err != nil {
		return &model.Customer{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (d *CustomerDatastore) GetAll(ctx *gofr.Context) ([]model.Customer, error) {
	var customers []model.Customer

	rows, err := ctx.DB().QueryContext(ctx, "SELECT  id, name, email, phone, address, city, date_of_birth, is_active FROM customers")
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.Address, &customer.City, &customer.DateOfBirth, &customer.IsActive); err != nil {
			return nil, errors.DB{Err: err}
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (d *CustomerDatastore) GetByID(ctx *gofr.Context, id string) (*model.Customer, error) {
	var resp model.Customer

	querySelect := "SELECT id, name, email, phone, address, city, date_of_birth, is_active FROM customers WHERE id = ?"

	err := ctx.DB().QueryRowContext(ctx, querySelect, id).Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Phone, &resp.Address, &resp.City, &resp.DateOfBirth, &resp.IsActive)

	switch err {
	case sql.ErrNoRows:
		return &model.Customer{}, errors.EntityNotFound{Entity: "customer", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Customer{}, errors.DB{Err: err}
	}
}

func (d *CustomerDatastore) Update(ctx *gofr.Context, customer *model.Customer) (*model.Customer, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET name = ?, email = ?, phone = ?, address = ?, city = ?, date_of_birth = ?, is_active = ? WHERE id = ?", customer.IsActive, customer.DateOfBirth, customer.City, customer.Address, customer.Phone, customer.Email, customer.Name, customer.ID)
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
