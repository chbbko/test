package usecases

import "github.com/test/entities"

type DoctorUseCase interface {
	List() ([]entities.Doctor,error)
}
