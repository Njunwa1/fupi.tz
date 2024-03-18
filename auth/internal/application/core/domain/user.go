package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Role struct {
	Name string
}

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Names       string
	Company     string
	Email       string
	PhoneNumber string
	Password    string
	Role        Role
	CreatedAt   time.Time
}

func NewUser(names, company, email, phoneNumber, password string, role Role) *User {
	return &User{
		Names:       names,
		Company:     company,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
		Role:        role,
	}
}

func (u *User) HashPassword(password string) (string, error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (u *User) VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
