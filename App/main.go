package main

import (
	"MyModule/App/handlers"
	"MyModule/App/model"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/coco"

	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Database Not Connected")
	}
	model.AutoMigrate(db)
	model.InitDefaultRoles(db)

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/app").Subrouter()

	handlers.RegisterHandlers(apiRouter, db)

	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", r)
}
