package repositories

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
)

// interface employee
type PhotoRepo interface {
	CreatePhoto(in models.Photo) (res models.Photo, err error)
	GetAllPhoto() (res []models.Photo, err error)
	GetPhotoById(id int64) (res models.Photo, err error)
	UpdatePhoto(in models.Photo) (res models.Photo, err error)
	DeletePhoto(id int64) (err error)
}

func (r Repo) CreatePhoto(in models.Photo) (res models.Photo, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllPhoto() (res []models.Photo, err error) {
	err = r.gorm.Preload("User").Preload("Comments").Find(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetPhotoById(id int64) (res models.Photo, err error) {
	err = r.gorm.Preload("User").Preload("Comments").First(&res, id).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdatePhoto(in models.Photo) (res models.Photo, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.Id).Updates(models.Photo{
		Title:    in.Title,
		Caption:  in.Caption,
		PhotoUrl: in.PhotoUrl,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeletePhoto(id int64) (err error) {
	err = r.gorm.Delete(&models.Photo{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
