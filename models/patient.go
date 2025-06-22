package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model `swaggerignore:"true"`
	Name       string
	Age        int
	Address    string
	Condition  string
	DoctorID   uint
	Username   string `gorm:"unique"`
	Password   string
	Role       string
}
