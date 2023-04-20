package repositories

import (
	"gorm.io/gorm"
)

type Repo struct {
	gorm *gorm.DB
}

type RepoInterface interface {
	UserRepo
	PhotoRepo
	CommentRepo
	SocialMediaRepo
}

// constructor function
func NewRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm}
}
