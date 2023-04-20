package models

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string       `json:"username" form:"username" gorm:"not null;uniqueIndex" valid:"required~username is required"`
	Email       string       `json:"email" form:"email" gorm:"not null;uniqueIndex" valid:"required~email is required,email~its must be a valid email"`
	Password    string       `json:"password" form:"password" gorm:"not null" valid:"required~password is required,minstringlength(6)~the length of password must be longer than 6 character"`
	Age         int          `json:"age" form:"age" gorm:"not null" valid:"required~age is required,range(8|100)~age must be greater than 8"`
	Photos      []Photo      `json:"photos,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments    []Comment    `json:"comments,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	SocialMedia *SocialMedia `json:"social_media,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helpers.HashPass(u.Password)

	return
}
