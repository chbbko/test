package models

import "github.com/test/entities"

type Patience struct {
	PinCode string `json:"pin_code"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type Schedule struct {
	Ref string
	Start int64
	End  int64
}

type AppointmentList struct {
	DoctorID string `json:"doctor_id"`
	Patience Patience `json:"patience"`
	Schedule Schedule `json:"details"`
}

type AppointmentCreate struct {
	PinCode string `json:"pin_code" validate:"required,alphanum"`
	Phone string `json:"phone" validate:"required,alphanum"`
	Ref string `json:"ref" validate:"required,alphanum"`
}

type AppointmentDelete struct {
	Ref string `json:"ref" validate:"required,alphanum"`
}

type AppointmentQueryString struct {
	DoctorID   string `form:"doctor_id" json:"doctor_id" validate:"required,alphanum"`
}

func ToModelAppointmentList(e []entities.AppointmentList) []AppointmentList {
	var m []AppointmentList
	for _, el := range e {
		cel := AppointmentList{}
		if err := copier.Copy(&cel, &el); err != nil {
			continue
		}
		m = append(m, cel)
	}
	return m
}

func ToEntitiesAppointmentCreate(m AppointmentCreate) (entities.AppointmentCreate, error) {
	var e entities.AppointmentCreate
	if err := copier.Copy(&e, &m); err != nil {
		return e, err
	}
	return e, nil
}
