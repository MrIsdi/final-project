package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Message string `gorm:"not null;type:varchar(255)" json:"message" form:"message" valid:"required~message must be required"`
	UserID  uint   `gorm:"foreignKey" json:"user_id" form:"user_id"`
	PhotoID uint   `gorm:"foreignKey" json:"photo_id" form:"photo_id"`
}

func (u *Comment) BeforeCreate(idb *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
