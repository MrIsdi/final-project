package models

import (
	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(255)" json:"username" form:"username" valid:"required~Username must be required"`
	Email    string `gorm:"not null;unique;type:varchar(255)" json:"email" form:"email" valid:"required~Email must be required,email~Invalid email format"`
	Password string `gorm:"not null;type:varchar(255)" json:"password" form:"password" valid:"required~Password must be required,minstringlength(6)~Password must be minimum 6 character"`
	Age      int64  `gorm:"not null;type:int(255)" json:"age" form:"age" valid:"required~Age must be required"`
}

func (u *User) BeforeCreate(idb *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
