package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null;type:varchar(255)" json:"title" form:"title" valid:"required~title must be required"`
	Caption  string `gorm:"not null;unique;type:varchar(255)" json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null;type:varchar(255)" json:"photo_url" form:"photo_url" valid:"required~photo_url must be required"`
	UserID   uint   `gorm:"foreignKey" json:"user_id" form:"user_id"`
}

func (u *Photo) BeforeCreate(idb *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
