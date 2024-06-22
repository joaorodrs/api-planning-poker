package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joaorodrs/api-planning-poker/database"
	"github.com/matoous/go-nanoid/v2"

	"github.com/joaorodrs/api-planning-poker/utils"
)

var db *gorm.DB

type User struct {
	ID        string     `gorm:"primaryKey"json:"id"`
	Name      string     `gorm:""json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
	Password  string     `json:"password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID, err = gonanoid.New()
	if err != nil {
		return err
	}

	return nil
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

func GetAllUsers(pagination *utils.Pagination) ([]User, *utils.Pagination) {
	var Users []User

	db.Scopes(utils.Paginate(&Users, pagination, db)).Find(&Users)

	return Users, pagination
}

func GetUserById(Id string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func UpdateUser(Id string, user User) User {
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

func DeleteUser(Id string) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}
