package models

import (
	"github.com/jinzhu/gorm"
)

//User Model
type User struct {
	gorm.Model
	ID    int
	Name  string
	Email string
}
