package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" gorm:"not null" valid:"required~name is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" gorm:"not null" valid:"required~social_media_url is required"`
	UserID         int    `json:"user_id" gorm:"not null;uniqueIndex"`
}

// hooks
func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return
}
