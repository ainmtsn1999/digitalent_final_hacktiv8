package services

import "github.com/ainmtsn1999/digitalent_final_hacktiv8/repositories"

type Service struct {
	repo repositories.RepoInterface
}

type ServiceInterface interface {
	UserService
	PhotoService
	CommentService
	SocialMediaService
}

func NewService(repo repositories.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
