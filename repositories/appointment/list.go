package appointment
import (
	"errors"
	"fmt"
	"github.com/test/entities"
	"github.com/test/repositories/models"
	"github.com/jinzhu/copier"
)

func (r *appointmentRepository) ListAppointment(params entities.AppointmentListParams) ([]entities.AppointmentList, error) {
	var m []models.Appointment
	var e []entities.AppointmentList
	//TODO:: preload booker details
	db := r.dbpg
	db = db.Model(&models.Appointment{})
	db = db.Select("id, doctor_id, ref, start, end, patience_pin_code")
	db = db.Where("UPPER(doctor_id) = ?", fmt.Sprintf("%s", params.DoctorID))
	//db = db.Limit(params.Limit).Offset(params.Skip).Order("created_at DESC").Order("ID DESC")

	result := db.Find(&m)
	fmt.Println("-----> result", m)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	copyError := copier.Copy(&e, m)

	if copyError != nil {
		return nil, copyError
	}
	return nil, nil
}
