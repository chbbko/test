package appointment
import (
	"errors"
	"github.com/test/entities"
	"github.com/test/repositories/models"

)

func (r *appointmentRepository) Cancel(params entities.AppointmentDelete) error {
	var m models.Appointment

	m.ID = params.ID
	db := r.dbpg
	db = db.Model(&models.Appointment{})
	result := db.Delete(&m)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("id not exist")
	}

	return nil
}
