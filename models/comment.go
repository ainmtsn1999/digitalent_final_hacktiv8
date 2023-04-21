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

type CommentResponse struct {
	Id      int            `json:"id"`
	Message string         `json:"message"`
	UserID  int            `json:"user_id,omitempty"`
	Photo   *PhotoResponse `json:"photo,omitempty"`
	User    *UserResponse  `json:"user,omitempty"`
}

func (u *Comment) ParseToModel() *CommentResponse {
	if u.User != nil {
		return &CommentResponse{
			Id:      u.Id,
			Message: u.Message,
			Photo:   u.Photo.ParseToModel(),
			User:    u.User.ParseToModel(),
		}
	} else {
		return &CommentResponse{
			Id:      u.Id,
			Message: u.Message,
			UserID:  u.UserID,
			Photo:   nil,
			User:    nil,
		}
	}
}

// hooks
func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return
}
