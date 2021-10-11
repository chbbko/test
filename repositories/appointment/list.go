package appointment
import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/test/entities"
	"github.com/test/repositories/models"
)

func (r *appointmentRepository) ListAppointment(params entities.AppointmentListParams) ([]entities.AppointmentList, error) {
	var m []models.Appointment
	var e []entities.AppointmentList
	//TODO:: preload booker details
	db := r.dbpg
	db = db.Model(&models.Appointment{})
	db = db.Select("id, doctor_id, ref, start, end, patience_pin_code")
	db = db.Where("UPPER(doctor_id) = ?", fmt.Sprintf("%s", params.DoctorID))

	result := db.Find(&m)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	copyError := copier.Copy(&e, m)

	if copyError != nil {
		return nil, copyError
	}
	return nil, nil
}
