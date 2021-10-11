package appointment

import "github.com/test/entities"

func (r appointmentUseCase) Create(params entities.AppointmentCreate) error{
	return r.repo.Create(params)
}
