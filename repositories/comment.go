package repositories

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
)

// interface employee
type CommentRepo interface {
	CreateComment(in models.Comment) (res models.Comment, err error)
	GetAllComment() (res []models.Comment, err error)
	GetAllCommentByUserId(id int64) (res []models.Comment, err error)
	GetCommentById(id int64) (res models.Comment, err error)
	UpdateComment(in models.Comment) (res models.Comment, err error)
	DeleteComment(id int64) (err error)
}

func (r Repo) CreateComment(in models.Comment) (res models.Comment, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllComment() (res []models.Comment, err error) {
	err = r.gorm.Preload("User").Preload("Photo").Find(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllCommentByUserId(id int64) (res []models.Comment, err error) {
	err = r.gorm.Where("user_id = ?", id).Preload("User").Preload("Photo").Find(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetCommentById(id int64) (res models.Comment, err error) {
	err = r.gorm.Preload("User").Preload("Photo").First(&res, id).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateComment(in models.Comment) (res models.Comment, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.Id).Updates(models.Comment{
		PhotoID: in.PhotoID,
		UserID:  in.UserID,
		Message: in.Message,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteComment(id int64) (err error) {
	err = r.gorm.Delete(&models.Comment{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
