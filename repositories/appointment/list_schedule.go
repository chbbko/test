package appointment
import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/test/entities"
	"github.com/test/repositories/models"
)

func (r *appointmentRepository) ListSchedule(params entities.ScheduleParams) ([]entities.ScheduleList, error) {
	var m []models.Schedule
	var e []entities.ScheduleList

	db := r.dbpg
	db = db.Model(&models.Schedule{})
	db = db.Select("*")
	if params.DoctorID != "" {
		db = db.Where("UPPER(doctor_id) = ?", fmt.Sprintf("%s", params.DoctorID))
	}
	//db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

	result := db.Find(&m)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	copyError := copier.Copy(&e, m)
	if copyError != nil {
		return nil, copyError
	}
	return e, nil
}
