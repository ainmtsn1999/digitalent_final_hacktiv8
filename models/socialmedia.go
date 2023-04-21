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

type SocialMediaResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
}

func (u *SocialMedia) ParseToModel() *SocialMediaResponse {
	return &SocialMediaResponse{
		Id:             u.Id,
		Name:           u.Name,
		SocialMediaUrl: u.SocialMediaUrl,
		UserID:         u.UserID,
	}
}

// hooks
func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return
}
