package entities

import "time"

type AppointmentList struct {
	PatiencePinCode string
	DoctorID string
	Phone string
	Ref string
	Start *time.Time
	End *time.Time
}

type AppointmentCreate struct {
	PatiencePinCode string
	DoctorID string
	Ref string
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
