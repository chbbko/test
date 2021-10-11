package doctorlist
import (
	"github.com/test/repositories"
	"github.com/test/usecases"
)

type doctorUseCase struct {
	repo repositories.DoctorRepository
}

func NewUseCase(repo repositories.DoctorRepository) usecases.DoctorUseCase {
	return doctorUseCase{repo: repo}
}
