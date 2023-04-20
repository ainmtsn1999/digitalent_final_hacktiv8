package services

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
)

// interface employee
type UserService interface {
	UserRegister(in models.User) (res models.User, err error)
	UserLogin(in models.User) (res models.User, err error)
}

func (s *Service) UserRegister(in models.User) (res models.User, err error) {
	return s.repo.UserRegister(in)
}
func (s *Service) UserLogin(in models.User) (res models.User, err error) {
	return s.repo.UserLogin(in)
}
