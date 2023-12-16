package datastore

import (
	"car-service/model"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func TestCoreLayer(*testing.T) {
	app := gofr.New()

	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true

	createTable(app)
}

func createTable(app *gofr.Gofr) {
	_, err := app.DB().Exec("DROP TABLE IF EXISTS customers;")

	if err != nil {
		return
	}

	_, err = app.DB().Exec("CREATE TABLE IF NOT EXISTS customers " +
		"(id serial primary key, name varchar(255), email varchar(255), phone varchar(255), address varchar(255), city varchar(255),date_of_birth varchar(255), is_active boolean);")
	if err != nil {
		return
	}
}

func TestAddCustomer(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()
	tests := []struct {
		desc     string
		customer model.Customer
		mockErr  error
		err      error
	}{
		{"Valid case", model.Customer{
			ID:          1,
			Name:        "Chetan Kushwah",
			Email:       "chetankushwah929@gmail.com",
			Phone:       "1234567890",
			Address:     "123 Main St",
			City:        "Anytown",
			DateOfBirth: "1990-01-01",
			IsActive:    true,
		}, nil, nil}, {
			"DB error", model.Customer{

				ID:          2,
				Name:        "Chetan Kush",
				Email:       "chetankushwah929@gmail.com",
				Phone:       "1234547890",
				Address:     "123 Main St",
				City:        "Anytown",
				DateOfBirth: "1910-01-01",
				IsActive:    true,
			}, errors.DB{}, errors.DB{Err: errors.DB{}}},
	}

	for i, tc := range tests {
		// Set up the expectations for the INSERT query
		mock.ExpectExec("INSERT INTO customers (id, name, email, phone, address, city, date_of_birth, is_active) VALUES (?,?,?,?,?,?,?,?)").
			WithArgs(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.Address, tc.customer.City, tc.customer.DateOfBirth, tc.customer.IsActive).
			WillReturnResult(sqlmock.NewResult(2, 1)).
			WillReturnError(tc.mockErr)

		// Set up the expectations for the SELECT query
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "city", "date_of_birth", "is_active"}).
			AddRow(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.Address, tc.customer.City, tc.customer.DateOfBirth, tc.customer.IsActive)
		mock.ExpectQuery("SELECT id, name, email, phone, address, city, date_of_birth, is_active FROM customers WHERE id = ?").
			WithArgs(tc.customer.ID).
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)

		datastore := NewCustomerDatastore()
		resp, err := datastore.Create(ctx, &tc.customer)

		ctx.Logger.Log(resp)
		assert.IsType(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}

}

func TestGetAllCustomers(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock co nnection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc      string
		customers []model.Customer
		mockErr   error
		err       error
	}{
		{"Valid case with customers", []model.Customer{
			{ID: 1, Name: "John Doe", Email: "chetankushwah929@gmail.com", Phone: "1234567890", City: "New York", DateOfBirth: "1990-01-01", IsActive: true},
			{ID: 2, Name: "Jane Smith", Email: "chetankushwah929@gmail.com", Phone: "9876543210", City: "San Francisco", DateOfBirth: "1990-01-01", IsActive: true},
		}, nil, nil},
		{"Valid case with no customers", []model.Customer{}, nil, nil},
		{"Error case", nil, errors.Error("database error"), errors.DB{Err: errors.Error("database error")}},
	}

	for i, tc := range tests {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "city", "date_of_birth", "is_active"})
		for _, customer := range tc.customers {
			rows.AddRow(customer.ID, customer.Name, customer.Email, customer.Phone, customer.Address, customer.City, customer.DateOfBirth, customer.IsActive)
		}

		mock.ExpectQuery("SELECT id, name, email, phone, address, city, date_of_birth, is_active FROM customers").
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)

		datastore := NewCustomerDatastore()
		resp, err := datastore.GetAll(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, tc.customers, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestGetCustomerByID(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc     string
		id       string
		customer model.Customer
		mockErr  error
		err      error
	}{
		{"Valid case", "1", model.Customer{ID: 1, Name: "Chetan Kushwah", Email: "chetankushwah929@gmail.com", Phone: "1234567890", City: "New York", DateOfBirth: "1990-01-01", IsActive: true}, nil, nil},
		{"Entity not found", "2", model.Customer{}, sql.ErrNoRows, errors.EntityNotFound{Entity: "customer", ID: "2"}},
		{"Error case", "3", model.Customer{}, errors.Error("database error"), errors.DB{Err: errors.Error("database error")}},
	}

	for i, tc := range tests {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "city", "date_of_birth", "is_active"}).
			AddRow(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.Address, tc.customer.City, tc.customer.DateOfBirth, tc.customer.IsActive)

		mock.ExpectQuery("SELECT id, name, email, phone, address, city, date_of_birth, is_active FROM customers WHERE id = ?").
			WithArgs(tc.id).
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)

		datastore := NewCustomerDatastore()
		resp, err := datastore.GetByID(ctx, tc.id)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, &tc.customer, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestUpdateCustomer(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc     string
		customer model.Customer
		mockErr  error
		err      error
	}{
		{"Valid case", model.Customer{ID: 1, Name: "Updated Name", Email: "updated@gmail.com", Phone: "9876543210", Address: "Updated Address", City: "Updated City", DateOfBirth: "1998-01-01", IsActive: true}, nil, nil},
		{"Entity not found", model.Customer{ID: 2, Name: "Updated Name", Email: "updated@gmail.com", Phone: "9876573210", Address: "Updated Address", City: "Updated City", DateOfBirth: "1998-01-01", IsActive: true}, sql.ErrNoRows, errors.EntityNotFound{Entity: "customer", ID: "2"}},
		{"Error case", model.Customer{ID: 3, Name: "Updated Name", Email: "updated@gmail.com", Phone: "9876543810", Address: "Updated Address", City: "Updated City", DateOfBirth: "1998-01-01", IsActive: true}, errors.Error("database error"), errors.DB{Err: errors.Error("database error")}},
	}

	for i, tc := range tests {
		mock.ExpectExec("UPDATE customers SET name = ?, email = ?, phone = ?, address = ?, city = ?, date_of_birth = ?, is_active = ? WHERE id = ?").
			WithArgs(tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.Address, tc.customer.City, tc.customer.DateOfBirth, tc.customer.IsActive, tc.customer.ID).
			WillReturnResult(sqlmock.NewResult(2, 1)).
			WillReturnError(tc.mockErr)

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "city", "date_of_birth", "is_active"}).
			AddRow(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.Address, tc.customer.City, tc.customer.DateOfBirth, tc.customer.IsActive)

		mock.ExpectQuery("SELECT id, name, email, phone, address, city, date_of_birth, is_active FROM customers WHERE id = ?").
			WithArgs(tc.customer.ID).
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)

		datastore := NewCustomerDatastore()
		resp, err := datastore.Update(ctx, &tc.customer)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, &tc.customer, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestDeleteCustomer(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	tests := []struct {
		desc    string
		id      string
		mockErr error
		err     error
	}{
		{"Valid case", "1", nil, nil},
		{"Entity not found", "2", sql.ErrNoRows, errors.EntityNotFound{Entity: "customer", ID: "2"}},
		{"Error case", "3", errors.Error("database error"), errors.DB{Err: errors.Error("database error")}},
	}

	for i, tc := range tests {
		mock.ExpectExec("DELETE FROM customers WHERE id = ?").
			WithArgs(tc.id).
			WillReturnResult(sqlmock.NewResult(2, 1)).
			WillReturnError(tc.mockErr)

		datastore := NewCustomerDatastore()
		err := datastore.Delete(ctx, tc.id)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
