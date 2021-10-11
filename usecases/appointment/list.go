package appointment

import "github.com/test/entities"

func (r appointmentUseCase) List(params entities.AppointmentListParams) ([]entities.AppointmentList,error) {
	return r.repo.ListAppointment(params)
}
