package doctors

import (
	"github.com/jinzhu/gorm"
	"github.com/test/repositories"
)

type doctorRepository struct {
	dbpg *gorm.DB
}

type DoctorRepositoryParams struct {
	DB *gorm.DB
}

func NewDoctorRepository(params DoctorRepositoryParams) repositories.DoctorRepository {
	return &doctorRepository{
		dbpg: params.DB,
	}
}
