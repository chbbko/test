package models
import (
	"github.com/jinzhu/gorm"
	"time"
)
type Appointment struct {
	gorm.Model
	PatiencePinCode string
	DoctorID string
	Phone string
	Ref string
	Start time.Time
	End time.Time
}
