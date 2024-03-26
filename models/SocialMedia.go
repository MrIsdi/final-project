package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"not null;type:varchar(255)" json:"name" form:"name" valid:"required~name must be required"`
	SocialMediaUrl string `gorm:"not null;type:varchar(255)" json:"social_media_url" form:"social_media_url" valid:"required~social_media_url must be required"`
	UserID         uint   `gorm:"foreignKey" json:"user_id" form:"user_id"`
}

func (u *SocialMedia) BeforeCreate(idb *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
