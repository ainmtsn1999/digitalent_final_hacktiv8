package services

import "github.com/ainmtsn1999/digitalent_final_hacktiv8/models"

type CommentService interface {
	CreateComment(in models.Comment) (res models.Comment, err error)
	GetAllComment() (res []models.Comment, err error)
	GetCommentById(id int64) (res models.Comment, err error)
	UpdateComment(in models.Comment) (res models.Comment, err error)
	DeleteComment(id int64) (err error)
}

func (s *Service) CreateComment(in models.Comment) (res models.Comment, err error) {
	return s.repo.CreateComment(in)
}
func (s *Service) GetAllComment() (res []models.Comment, err error) {
	return s.repo.GetAllComment()
}
func (s *Service) GetCommentById(id int64) (res models.Comment, err error) {
	return s.repo.GetCommentById(id)
}
func (s *Service) UpdateComment(in models.Comment) (res models.Comment, err error) {
	return s.repo.UpdateComment(in)
}
func (s *Service) DeleteComment(id int64) (err error) {
	return s.repo.DeleteComment(id)
}
