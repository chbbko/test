package doctors

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/test/entities"
	"github.com/test/repositories/models"
)

func (r *doctorRepository) ListDoctor() ([]entities.Doctor, error) {
	var m []models.Doctor
	var e []entities.Doctor

	db := r.dbpg
	db = db.Model(&models.Doctor{})
	db = db.Select("id, created_at, doctor_id, name")

	result := db.Find(&m)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	copyError := copier.Copy(&e, m)

	if copyError != nil {
		return nil, copyError
	}
	return e,nil
}
