package main

import (
	"car-service/datastore"
	"car-service/handler"

	"gofr.dev/pkg/gofr"
)

func main(){
	app := gofr.New()
	
	datastore := datastore.NewCustomerDatastore()
	handler := handler.NewHandler(datastore)

	app.POST("/customer", handler.Create)
	app.GET("/customer", handler.GetAll)
	app.GET("/customer/{id}",handler.GetByID)
    app.PUT("/customer",handler.Update)
    app.DELETE("/customer/{id}",handler.Delete)
	
	app.Start()

}