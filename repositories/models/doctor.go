package models

import (
	"github.com/jinzhu/gorm"
)

type Doctor struct {
	gorm.Model
	DoctorID string
	Name string
}

type BookingHistory struct {
	PinCode string
	Name string
}

