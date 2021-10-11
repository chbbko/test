package usecases
import "github.com/test/entities"
type AppointmentUseCase interface {
	List(entities.AppointmentListParams) ([]entities.AppointmentList,error)
	Cancel(entities.AppointmentDelete) error
	Create(entities.AppointmentCreate) error
}
