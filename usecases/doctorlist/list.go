package doctorlist

import "github.com/test/entities"

func (r doctorUseCase) List() ([]entities.Doctor,error) {
	return r.repo.ListDoctor()
}
