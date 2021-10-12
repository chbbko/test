package entities

import (
	"time"
)

type AppointmentList struct {
	PatiencePinCode string
	DoctorID string
	Phone string
	Ref string
	Start time.Time
	End time.Time
	ID uint
}

type AppointmentCreate struct {
	PatiencePinCode string
	DoctorID string
	Ref string
	Phone string
	Start time.Time
	End time.Time
}

type AppointmentDelete struct {
	ID uint
	Ref string
}

type AppointmentListParams struct {
	DoctorID string
}

type Doctor struct {
	DoctorID string
	Name string
}

type ScheduleParams struct {
	DoctorID   string
	Begin *int64
	End *int64
}

type ScheduleList struct {
	ID uint
	Ref string
	DoctorID string
	Start *time.Time
	End  *time.Time
}
