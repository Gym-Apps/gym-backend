package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountType int8
type Gender int8

const (
	Teacher AccountType = iota + 1
	Athlete
)

const (
	Woman Gender = iota + 1
	Man
)

type User struct {
	gorm.Model
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Phone       string      `json:"phone"`
	Password    string      `json:"password"`
	Email       string      `json:"email"`
	Birthday    time.Time   `json:"birthday"`
	AccountType AccountType `json:"account_type"`
	AccountName string      `json:"account_name" gorm:"-"`
	Gender      Gender      `json:"gender"`
	GenderName  string      `json:"gender_name" gorm:"-"`
}

func (u *User) AfterFind(db *gorm.DB) error {
	switch u.AccountType {
	case Teacher:
		u.AccountName = "Eğitmen"
	case Athlete:
		u.AccountName = "Sporcu"
	}

	switch u.Gender {
	case Woman:
		u.GenderName = "Kadın"
	case Man:
		u.GenderName = "Erkek"
	}
	return nil
}
