package services

import "github.com/ainmtsn1999/digitalent_final_hacktiv8/models"

type SocialMediaService interface {
	CreateSocialMedia(in models.SocialMedia) (res models.SocialMediaResponse, err error)
	GetAllSocialMedia() (res []models.SocialMediaResponse, err error)
	GetSocialMediaById(id int64) (res models.SocialMediaResponse, err error)
	GetSocialMediaByUserId(id int64) (res models.SocialMediaResponse, err error)
	UpdateSocialMedia(in models.SocialMedia) (res models.SocialMediaResponse, err error)
	DeleteSocialMedia(id int64) (err error)
}

func (s *Service) CreateSocialMedia(in models.SocialMedia) (result models.SocialMediaResponse, err error) {
	res, err := s.repo.CreateSocialMedia(in)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) GetAllSocialMedia() (result []models.SocialMediaResponse, err error) {
	res, err := s.repo.GetAllSocialMedia()
	if err != nil {
		return result, err
	}

	for _, sosmed := range res {
		result = append(result, *sosmed.ParseToModel())
	}

	return result, nil
}
func (s *Service) GetSocialMediaById(id int64) (result models.SocialMediaResponse, err error) {
	res, err := s.repo.GetSocialMediaById(id)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) GetSocialMediaByUserId(id int64) (result models.SocialMediaResponse, err error) {
	res, err := s.repo.GetSocialMediaByUserId(id)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) UpdateSocialMedia(in models.SocialMedia) (result models.SocialMediaResponse, err error) {
	res, err := s.repo.UpdateSocialMedia(in)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) DeleteSocialMedia(id int64) (err error) {
	return s.repo.DeleteSocialMedia(id)
}
