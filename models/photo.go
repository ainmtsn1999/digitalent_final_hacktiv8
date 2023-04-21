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

type PhotoResponse struct {
	Id       int               `json:"id"`
	Title    string            `json:"title"`
	Caption  string            `json:"caption"`
	PhotoUrl string            `json:"photo_url"`
	UserID   int               `json:"user_id,omitempty"`
	User     *UserResponse     `json:"user,omitempty"`
	Comments []CommentResponse `json:"comments,omitempty"`
}

func (u *Photo) ParseToModel() *PhotoResponse {
	//get all comment
	commentResp := []CommentResponse{}
	for _, comment := range u.Comments {
		commentResp = append(commentResp, *comment.ParseToModel())
	}

	if u.User != nil {
		return &PhotoResponse{
			Id:       u.Id,
			Title:    u.Title,
			Caption:  u.Caption,
			PhotoUrl: u.PhotoUrl,
			User:     u.User.ParseToModel(),
			Comments: commentResp,
		}
	} else {
		return &PhotoResponse{
			Id:       u.Id,
			Title:    u.Title,
			Caption:  u.Caption,
			PhotoUrl: u.PhotoUrl,
			UserID:   u.UserID,
			User:     nil,
			Comments: commentResp,
		}
	}

}

// hooks
func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return
}
