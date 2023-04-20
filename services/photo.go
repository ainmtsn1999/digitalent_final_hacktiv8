package services

import "github.com/ainmtsn1999/digitalent_final_hacktiv8/models"

type PhotoService interface {
	CreatePhoto(in models.Photo) (res models.Photo, err error)
	GetAllPhoto() (res []models.Photo, err error)
	GetPhotoById(id int64) (res models.Photo, err error)
	UpdatePhoto(in models.Photo) (res models.Photo, err error)
	DeletePhoto(id int64) (err error)
}

func (s *Service) CreatePhoto(in models.Photo) (res models.Photo, err error) {
	return s.repo.CreatePhoto(in)
}
func (s *Service) GetAllPhoto() (res []models.Photo, err error) {
	return s.repo.GetAllPhoto()
}
func (s *Service) GetPhotoById(id int64) (res models.Photo, err error) {
	return s.repo.GetPhotoById(id)
}
func (s *Service) UpdatePhoto(in models.Photo) (res models.Photo, err error) {
	return s.repo.UpdatePhoto(in)
}
func (s *Service) DeletePhoto(id int64) (err error) {
	return s.repo.DeletePhoto(id)
}
