package datastore

import (
	"car-service/model"
	"context"
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
		customer model.customer
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
			WithArgs(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.Address, tc.customer.City, tc.customer.Date_Of_Birth, tc.customer.Is_Active).
			WillReturnResult(sqlmock.NewResult(2, 1)).
			WillReturnError(tc.mockErr)

		// Set up the expectations for the SELECT query
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "city", "date_of_birth", "is_active"}).
			AddRow(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.Address, tc.customer.City, tc.customer.Date_Of_Birth, tc.customer.Is_Active)
		mock.ExpectQuery("SELECT id, name, email, phone, address, city, date_of_birth, is_active FROM employees WHERE id = ?").
			WithArgs(tc.customer.ID).
			WillReturnRows(rows).
			WillReturnError(tc.mockErr)

		store := New()
		resp, err := store.Create(ctx, tc.customer)

		ctx.Logger.Log(resp)
		assert.IsType(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}

}
