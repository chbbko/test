package models
import (
	"github.com/jinzhu/gorm"
)
type Patience struct {
	gorm.Model
	PinCode string
	Name string
	Phone string
}
