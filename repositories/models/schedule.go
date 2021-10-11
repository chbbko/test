package models

import (
	"time"
	"github.com/jinzhu/gorm"

)

type Schedule struct {
	gorm.Model
	Ref string
	DoctorID string
	Start *time.Time
	End  *time.Time
	IsCancelled bool
	BookerID *uint  `sql:"type:integer REFERENCES patiences(id)"`
}
