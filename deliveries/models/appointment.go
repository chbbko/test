package models

import (
	"github.com/jinzhu/copier"
	"github.com/test/entities"
	"time"
)

type Patience struct {
	PinCode string `json:"pin_code"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type Schedule struct {
	DoctorID string
	Ref string
	Start int64
	End  int64
}

type AppointmentList struct {
	DoctorID string `json:"doctor_id"`
	PatiencePinCode string `json:"pin_code"`
	Ref string `json:"ref"`
	Start int64 `json:"start"`
	End int64 `json:"end"`
	ID uint `json:"id"`
	//Patience Patience `json:"patience"`
	//Schedule Schedule `json:"details"`
}

type AppointmentCreate struct {
	PatiencePinCode string `json:"pin_code" validate:"required,alphanum"`
	Phone string `json:"phone" validate:"required,alphanum"`
	Ref string `json:"ref" validate:"required,alphanum"`
	Start int64 `json:"start" validate:"required"`
	End int64 `json:"end" validate:"required"`
	DoctorID string `json:"doctor_id"`
}

type AppointmentDelete struct {
	Ref string `json:"ref" validate:"required,alphanum"`
}

type AppointmentQueryString struct {
	DoctorID   *string `form:"doctor_id" json:"doctor_id" validate:"omitempty,alphanum"`
}

type ScheduleQueryString struct {
	DoctorID   *string `form:"doctor_id" json:"doctor_id" validate:"omitempty,alphanum"`
	Begin *int64 `form:"begin" json:"begin" validate:"omitempty"`
	End *int64 `form:"end" json:"end" validate:"omitempty"`
}

func ToModelAppointmentList(e []entities.AppointmentList) []AppointmentList {
	var m []AppointmentList
	for _, el := range e {
		cel := AppointmentList{}
		if err := copier.Copy(&cel, el); err != nil {
			continue
		}
		cel.Start = el.Start.UnixNano()
		cel.End = el.End.UnixNano()
		m = append(m, cel)
	}
	return m
}

func ToEntitiesAppointmentCreate(m AppointmentCreate) (entities.AppointmentCreate, error) {
	var e entities.AppointmentCreate
	if err := copier.Copy(&e, &m); err != nil {
		return e, err
	}
	e.Start = time.Unix(0, m.Start)
	e.End = time.Unix(0, m.End)
	return e, nil
}

func ToModelScheduleList(e []entities.ScheduleList) []Schedule {
	var m []Schedule
	for _, el := range e {
		cel := Schedule{}
		if err := copier.Copy(&cel, el); err != nil {
			continue
		}
		cel.Start = el.Start.UnixNano()
		cel.End = el.End.UnixNano()
		m = append(m, cel)
	}
	return m
}
