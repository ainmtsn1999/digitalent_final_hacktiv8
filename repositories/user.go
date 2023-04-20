package repositories

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
)

// interface employee
type UserRepo interface {
	UserRegister(in models.User) (res models.User, err error)
	UserLogin(in models.User) (res models.User, err error)
}

func (r Repo) UserRegister(in models.User) (res models.User, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UserLogin(in models.User) (res models.User, err error) {

	if in.Email != "" {
		err = r.gorm.Where("email = ? ", in.Email).First(&res).Error

		if err != nil {
			return res, err
		}
	} else if in.Username != "" {
		err = r.gorm.Where("username = ? ", in.Username).First(&res).Error

		if err != nil {
			return res, err
		}
	}

	return res, nil
}
