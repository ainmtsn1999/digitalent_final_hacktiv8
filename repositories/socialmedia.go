package repositories

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
)

// interface employee
type SocialMediaRepo interface {
	CreateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	GetAllSocialMedia() (res []models.SocialMedia, err error)
	GetSocialMediaById(id int64) (res models.SocialMedia, err error)
	GetSocialMediaByUserId(id int64) (res models.SocialMedia, err error)
	UpdateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	DeleteSocialMedia(id int64) (err error)
}

func (r Repo) CreateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllSocialMedia() (res []models.SocialMedia, err error) {
	err = r.gorm.Find(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetSocialMediaById(id int64) (res models.SocialMedia, err error) {
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetSocialMediaByUserId(id int64) (res models.SocialMedia, err error) {
	err = r.gorm.Where("user_id = ?", id).First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.Id).Updates(models.SocialMedia{
		Name:           in.Name,
		SocialMediaUrl: in.SocialMediaUrl,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteSocialMedia(id int64) (err error) {
	err = r.gorm.Delete(&models.SocialMedia{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
