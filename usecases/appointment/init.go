package appointment
import (
	"github.com/test/repositories"
	"github.com/test/usecases"
)

type appointmentUseCase struct {
	repo repositories.AppointmentRepository
}

func NewUseCase(repo repositories.AppointmentRepository) usecases.AppointmentUseCase {
	return appointmentUseCase{repo: repo}
}
