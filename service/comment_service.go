package service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helpers"
	"final-project/repository/comment_repository"
	"net/http"
)

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

type CommentService interface {
	CreateComment(newComment dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	GetComments() (*dto.CommentResultResponse, errs.MessageErr)
	GetCommentById(commentId int) (*dto.CommentResultOneResponse, errs.MessageErr)
	UpdateComment(commentId int, userId int, photoId int, payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	DeleteComment(commentId int) (*dto.NewCommentResponse, errs.MessageErr)
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}
func (commentService *commentService) CreateComment(newComment dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newComment)

	if err != nil {
		return nil, err
	}

	payload := entity.Comment{
		Message: newComment.Message,
		UserId:  newComment.UserId,
		PhotoId: newComment.PhotoId,
	}
	err = commentService.commentRepo.CreateComment(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "comment created",
	}
	return &response, nil

}
func (commentService *commentService) GetComments() (*dto.CommentResultResponse, errs.MessageErr) {
	comments, err := commentService.commentRepo.GetComments()
	if err != nil {
		return nil, err
	}

	var commentResponses []dto.CommentResponse
	for _, eachComment := range comments {
		commentResponses = append(commentResponses, eachComment.CommentEntityToDto())
	}

	result := dto.CommentResultResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment successfully to send",
		Data:       commentResponses,
	}
	return &result, nil
}
func (commentService *commentService) GetCommentById(commentId int) (*dto.CommentResultOneResponse, errs.MessageErr) {
	comment, err := commentService.commentRepo.GetCommentById(commentId)
	if err != nil {
		return nil, err
	}

	result := dto.CommentResultOneResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment successfully to send",
		Data:       comment.CommentEntityToDto(),
	}
	return &result, nil
}
func (commentService *commentService) UpdateComment(commentId int, userId int, photoId int, payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {
	comment := entity.Comment{
		Id:      commentId,
		UserId:  userId,
		PhotoId: photoId,
		Message: payload.Message,
	}
	_, err := commentService.commentRepo.GetCommentById(commentId)
	if err != nil {
		return nil, errs.NewNotFoundError("comment is not exists")
	}
	err = commentService.commentRepo.UpdateComment(comment)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	result := dto.NewCommentResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment successfully to update",
	}
	return &result, nil
}
func (commentService *commentService) DeleteComment(commentId int) (*dto.NewCommentResponse, errs.MessageErr) {
	_, err := commentService.commentRepo.GetCommentById(commentId)
	if err != nil {
		return nil, errs.NewNotFoundError("comment is not exists")
	}

	err = commentService.commentRepo.DeleteComment(commentId)
	if err != nil {
		return nil, errs.NewBadRequest("something went wrong")
	}

	result := dto.NewCommentResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment successfully to delete",
	}
	return &result, nil
}
