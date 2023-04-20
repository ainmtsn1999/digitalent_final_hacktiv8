package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Photo   *Photo `json:"photo,omitempty"`
	User    *User  `json:"user,omitempty"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
	Message string `json:"message" form:"message" gorm:"not null" valid:"required~title is required"`
}

// hooks
func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return
}
