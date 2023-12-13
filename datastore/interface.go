package datastore

import (
	"car-service/model"
	"gofr.dev/pkg/gofr"
)


type Customer interface {

	Create(ctx *gofr.Context, customer *model.Customer) (*model.Customer, error)

	GetAll(ctx *gofr.Context) ([]model.Customer, error)
}
