package datastore

import (
	"car-service/model"
	"gofr.dev/pkg/gofr"
)


type Customer interface {

	Create(ctx *gofr.Context, customer *model.Customer) (*model.Customer, error)
    
	GetAll(ctx *gofr.Context) ([]model.Customer, error)

    GetByID(ctx *gofr.Context, id string) (*model.Customer,error)

	Update(ctx *gofr.Context, customer *model.Customer) (*model.Customer, error)

	Delete(ctx *gofr.Context, id string) error
}
