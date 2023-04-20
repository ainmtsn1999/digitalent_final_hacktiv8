package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `json:"title" form:"title" gorm:"not null" valid:"required~title is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoUrl string    `json:"photo_url" form:"photo_url" gorm:"not null" valid:"required~photo is required"`
	UserID   int       `json:"user_id"`
	User     *User     `json:"user,omitempty"`
	Comments []Comment `json:"comments,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// hooks
func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return
}
