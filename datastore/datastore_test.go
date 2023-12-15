package datastore

import (
	"testing"

	"gofr.dev/pkg/datastore"
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
