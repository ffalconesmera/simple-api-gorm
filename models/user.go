package models

import (
	"errors"
	"fmt"
)

var Genders = map[int]string{1: "Hombre", 2: "Mujer"}

type User struct {
	Name      string `json:"name" gorm:"index;not null"`
	Email     string `json:"email" gorm:"index;not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Date      string `json:"date"`
	Gender    int    `json:"gender" binding:"required"`
	Password  string `json:"password"`
	Model
}

func (s *User) ValidateBase() error {
	if s == nil {
		return EmptyMesage("user")
	}

	errName := EmptyString(s.Name, "name")
	if errName != nil {
		return errName
	}

	errEmail := EmptyString(s.Email, "email")
	if errEmail != nil {
		return errEmail
	}

	errFirstName := EmptyString(s.FirstName, "first name")
	if errFirstName != nil {
		return errFirstName
	}

	errLastName := EmptyString(s.LastName, "Last name")
	if errLastName != nil {
		return errLastName
	}

	errDate := DateInvalid(s.Date, "date")
	if errDate != nil {
		return errDate
	}

	_, okGender := Genders[s.Gender]

	if !okGender {
		err := fmt.Sprintf("gender: is not a valid value. %v", Genders)
		return errors.New(err)
	}

	return nil
}

func (s *User) ValidateInsert() error {
	evi := s.ValidateBase()
	if evi != nil {
		return evi
	}

	errPassword := PasswordInvalid(s.Password, "password")
	if errPassword != nil {
		return errPassword
	}

	return nil
}
