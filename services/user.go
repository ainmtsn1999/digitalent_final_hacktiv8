package services

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
)

// interface employee
type UserService interface {
	UserRegister(in models.User) (res models.UserResponse, err error)
	UserLogin(in models.User) (res models.User, err error)
}

func (s *Service) UserRegister(in models.User) (result models.UserResponse, err error) {
	res, err := s.repo.UserRegister(in)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}

func (s *Service) UserLogin(in models.User) (res models.User, err error) {
	return s.repo.UserLogin(in)
}
