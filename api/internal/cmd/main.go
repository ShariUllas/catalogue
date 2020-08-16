package main

import (
	"catalogue/api/internal/app"
	"catalogue/api/internal/config"
	"catalogue/api/internal/core/data"
	"catalogue/api/internal/core/service"
	"flag"

	"log"
	"net/http"
)

var migDir = "./internal/core/data/migrations"

func main() {

	serve := flag.Bool("serve", false, "a bool")
	flag.Parse()
	if *serve {
		handler := setUpServer()
		router := handler.InitRouter()
		log.Fatal(http.ListenAndServe(handler.Port, router))
	}

}

func setUpServer() *app.Handler {
	handler := new(app.Handler)
	conf := config.NewConfig()
	handler.Config = conf
	dbConn := handler.GetDBConnectionString()
	db, err := config.GetDBConnection(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	config.MigrateFolder(dbConn, migDir)
	categoryRepo := data.NewCategoryRepo(db)
	productRepo := data.NewProductRepo(db)
	variantRepo := data.NewVariantRepo(db)
	handler.Category = service.NewCategoryService(categoryRepo)
	handler.Product = service.NewProductService(productRepo)
	handler.Variant = service.NewVariantService(variantRepo)
	return handler
}
