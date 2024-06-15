package models

import (
	"github.com/jinzhu/gorm"
	"github.com/joaorodrs/api-planning-poker/database"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `gorm:""json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&User{})
}

func (user *User) CreateUser() *User {
	db.NewRecord(user)
	db.Create(&user)
	return user
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func UpdateUser(Id int64, user User) User {
	var userDetails User
	db.Where("ID=?", Id).Find(&userDetails)

	if user.Name != "" {
		userDetails.Name = user.Name
	}
	if user.Email != "" {
		userDetails.Email = user.Email
	}

	db.Save(&userDetails)
	return userDetails
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}
