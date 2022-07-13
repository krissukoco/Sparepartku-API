package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Person struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique_id"`
	Password  string         `json:"-"`
	Phone     string         `json:"phone"`
	ImageURL  string         `json:"image_url"`
	Location  string         `json:"location"`
}

func NewUser(name, email, password, phone, location string) Person {
	// TODO: HASH Password
	id := uuid.New().String()
	pwdByte := []byte(password)
	hashedPwd, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("ERROR: Generating Hashed Password: ", err.Error())
	}
	user := Person{ID: id, Name: name, Email: email, Password: string(hashedPwd), Phone: phone, Location: location}
	return user
}
