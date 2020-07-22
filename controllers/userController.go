package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryansat/rest/models"
)

func GetUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:root@/db_lokasi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []models.User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")
	db, err := gorm.Open("mysql", "root:root@/db_lokasi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// vars := mux.Vars(r)
	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Println("name :" + name)
	fmt.Println("email :" + email)

	db.Create(&models.User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:root@/db_lokasi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// vars := mux.Vars(r)
	name := r.FormValue("name")

	var user models.User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:root@/db_lokasi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")

	var user models.User
	db.Where("id = ?", id).Find(&user)

	user.Name = name
	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func GetSingleUserEndPoint(w http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open("mysql", "root:root@/db_lokasi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	name := req.FormValue("name")

	var users []models.User
	db.Where("name = ?", name).Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func InitialMigration() {
	db, err := gorm.Open("mysql", "root:root@/db_lokasi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}
