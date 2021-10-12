package repositories
import "github.com/test/entities"


type AppointmentRepository interface {
	ListAppointment(entities.AppointmentListParams) ([]entities.AppointmentList, error)
	Create(entities.AppointmentCreate) error
	Cancel(entities.AppointmentDelete) error
	ListSchedule(params entities.ScheduleParams) ([]entities.ScheduleList, error)
}
