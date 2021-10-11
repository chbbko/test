package repositories

import "github.com/test/entities"


type DoctorRepository interface {
	ListDoctor() ([]entities.Doctor, error)
}
