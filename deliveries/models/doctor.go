package models

import (

	"github.com/jinzhu/copier"
	"github.com/test/entities"

)

type DoctorList struct {
	ID         uint   `json:"id"`
	DoctorID string `json:"doctor_id"`
	Name string `json:"name"`
}

func ToModelDoctorList(e []entities.Doctor) []DoctorList {
	var m []DoctorList
	for _, el := range e {
		cel := DoctorList{}
		if err := copier.Copy(&cel, &el); err != nil {
			continue
		}
		m = append(m, cel)
	}
	return m
}
