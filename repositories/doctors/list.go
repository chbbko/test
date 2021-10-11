package doctors

import (
	"errors"
	"fmt"
	"github.com/test/entities"
	"github.com/test/repositories/models"
	"github.com/jinzhu/gorm"
)

func (r *doctorRepository) ListDoctor() ([]entities.Doctor, error) {
	var m []models.Doctor
	var e []entities.Doctor

	db := r.dbpg
	db = db.Model(&models.Doctor{})
	db = db.Select("id, created_at, doctor_id, name")
	//db = db.Where("UPPER(employee_id) = ?", fmt.Sprintf("%s", params.EmployeeID))
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
	return e,nil
}
