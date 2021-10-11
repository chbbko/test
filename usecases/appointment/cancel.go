package appointment
import "github.com/test/entities"

func (r appointmentUseCase) Cancel(params entities.AppointmentDelete) error{
	return r.repo.Cancel(params)
}
