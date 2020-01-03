package application

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ApplicationModel struct {
	gorm.Model
	Name  string `json:"name"`
	Token string `json:"token" gorm:"unique_index;not null"`
}
