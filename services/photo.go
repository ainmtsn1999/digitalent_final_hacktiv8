package services

import "github.com/ainmtsn1999/digitalent_final_hacktiv8/models"

type PhotoService interface {
	CreatePhoto(in models.Photo) (res models.PhotoResponse, err error)
	GetAllPhoto() (res []models.PhotoResponse, err error)
	GetAllPhotoByUserId(id int64) (res []models.PhotoResponse, err error)
	GetPhotoById(id int64) (res models.PhotoResponse, err error)
	UpdatePhoto(in models.Photo) (res models.PhotoResponse, err error)
	DeletePhoto(id int64) (err error)
}

func (s *Service) CreatePhoto(in models.Photo) (result models.PhotoResponse, err error) {
	res, err := s.repo.CreatePhoto(in)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) GetAllPhoto() (result []models.PhotoResponse, err error) {
	res, err := s.repo.GetAllPhoto()
	if err != nil {
		return result, err
	}

	for _, photo := range res {
		result = append(result, *photo.ParseToModel())
	}

	return result, nil
}
func (s *Service) GetAllPhotoByUserId(id int64) (result []models.PhotoResponse, err error) {
	res, err := s.repo.GetAllPhotoByUserId(id)
	if err != nil {
		return result, err
	}

	for _, photo := range res {
		result = append(result, *photo.ParseToModel())
	}

	return result, nil
}
func (s *Service) GetPhotoById(id int64) (result models.PhotoResponse, err error) {
	res, err := s.repo.GetPhotoById(id)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) UpdatePhoto(in models.Photo) (result models.PhotoResponse, err error) {
	res, err := s.repo.UpdatePhoto(in)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) DeletePhoto(id int64) (err error) {
	return s.repo.DeletePhoto(id)
}
