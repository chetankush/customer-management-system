package handler

import (
	"car-service/datastore"
	"car-service/model"
	"strconv"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	store datastore.Customer
}

func NewHandler(s datastore.Customer) Handler{
	return Handler{store: s}
}


func (h Handler) Create(ctx *gofr.Context) (interface{},error){
	var customer model.Customer

	if err := ctx.Bind(&customer); err != nil {
		ctx.Logger.Errorf("error in bindng: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	resp, err := h.store.Create(ctx, &customer)
	if err != nil{
		return nil, err
	}
	return resp, nil
}

func (h Handler) GetAll(ctx *gofr.Context) (interface{},error){
	resp, err := h.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}


func validateID(id string) (int, error){
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return res, err
}