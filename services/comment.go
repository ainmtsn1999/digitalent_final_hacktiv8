package services

import "github.com/ainmtsn1999/digitalent_final_hacktiv8/models"

type CommentService interface {
	CreateComment(in models.Comment) (res models.CommentResponse, err error)
	GetAllComment() (res []models.CommentResponse, err error)
	GetAllCommentByUserId(id int64) (res []models.CommentResponse, err error)
	GetCommentById(id int64) (res models.CommentResponse, err error)
	UpdateComment(in models.Comment) (res models.CommentResponse, err error)
	DeleteComment(id int64) (err error)
}

func (s *Service) CreateComment(in models.Comment) (result models.CommentResponse, err error) {
	res, err := s.repo.CreateComment(in)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) GetAllComment() (result []models.CommentResponse, err error) {
	res, err := s.repo.GetAllComment()
	if err != nil {
		return result, err
	}

	for _, comment := range res {
		result = append(result, *comment.ParseToModel())
	}

	return result, nil
}
func (s *Service) GetAllCommentByUserId(id int64) (result []models.CommentResponse, err error) {
	res, err := s.repo.GetAllCommentByUserId(id)
	if err != nil {
		return result, err
	}

	for _, comment := range res {
		result = append(result, *comment.ParseToModel())
	}

	return result, nil
}
func (s *Service) GetCommentById(id int64) (result models.CommentResponse, err error) {
	res, err := s.repo.GetCommentById(id)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) UpdateComment(in models.Comment) (result models.CommentResponse, err error) {
	res, err := s.repo.UpdateComment(in)
	if err != nil {
		return result, err
	}

	result = *res.ParseToModel()

	return result, nil
}
func (s *Service) DeleteComment(id int64) (err error) {
	return s.repo.DeleteComment(id)
}
