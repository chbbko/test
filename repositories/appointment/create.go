package appointment
import (
	"errors"
	"github.com/test/entities"
	"github.com/test/repositories/models"
	"github.com/jinzhu/copier"
)

func (r *appointmentRepository) Create(params entities.AppointmentCreate) error {
	var m models.Appointment
	err := copier.Copy(&m, params)
	if err != nil {
		return errors.New("parse data failed")
	}
	result := r.dbpg.Create(&m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
