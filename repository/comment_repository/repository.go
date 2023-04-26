package comment_repository

import (
	"final-project/entity"
	"final-project/pkg/errs"
)

type CommentRepository interface {
	CreateComment(newComment entity.Comment) errs.MessageErr
	GetComments() ([]*entity.Comment, errs.MessageErr)
	GetCommentById(commentId int) (*entity.Comment, errs.MessageErr)
	UpdateComment(payload entity.Comment) errs.MessageErr
	DeleteComment(commentId int) errs.MessageErr
}
