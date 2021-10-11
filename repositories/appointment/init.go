package appointment
import (
	"github.com/jinzhu/gorm"
	"github.com/test/repositories"
)

type appointmentRepository struct {
	dbpg *gorm.DB
}

type ParamsAppointmentRepository struct {
	DB *gorm.DB
}

func NewRepository(params ParamsAppointmentRepository) repositories.AppointmentRepository {
	return &appointmentRepository{
		dbpg: params.DB,
	}
}
