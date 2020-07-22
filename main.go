package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryansat/rest/controllers"
	"github.com/ryansat/rest/models"
)

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetUsersEndPoint).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUserEndPoint).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUserEndPoint).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUserEndPoint).Methods("PUT")
	router.HandleFunc("/api/users/{name}", controllers.GetSingleUserEndPoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func initialMigration() {
	db, err := gorm.Open("mysql", "root:root@/db_lokasi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")

	initialMigration()
	// Handle Subsequent requests
	handleRequests()
}
