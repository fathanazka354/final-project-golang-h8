package comment_pg

import (
	"database/sql"
	"errors"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/comment_repository"
)

const (
	createCommentQuery = `
		INSERT INTO "comments"
		(
			message,
			userid,
			photoid
		)
		VALUES ($1, $2, $3)
	`
	getCommentsQuery = `
		SELECT "id", "userid", "photoid", "message", "createdat", "updatedat" FROM "comments" 
		ORDER BY "id" ASC
	`
	getCommentByIdQuery = `
	SELECT "id", "userid", "photoid", "message", "createdat", "updatedat" FROM "comments" 
	WHERE "id" = $1
	`

	updateCommentQuery = `
	UPDATE "comments" SET "message" = $1
	WHERE "id" = $2
	`

	deleteCommentQuery = `
	DELETE FROM "comments" 
	WHERE "id" = $1
	`
)

type commentPG struct {
	db *sql.DB
}

func NewCommentPG(db *sql.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

// CreateComment(newComment entity.Comment) errs.MessageErr
// GetComments() ([]*entity.Comment, errs.MessageErr)
// GetCommentById() (*entity.Comment, errs.MessageErr)
// UpdateComment(payload entity.Comment) errs.MessageErr
// DeleteComment(commentId int) errs.MessageErr
func (commentPG *commentPG) CreateComment(newComment entity.Comment) errs.MessageErr {
	_, err := commentPG.db.Exec(createCommentQuery, newComment.Message, newComment.UserId, newComment.PhotoId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (commentPG *commentPG) GetComments() ([]*entity.Comment, errs.MessageErr) {
	responses, err := commentPG.db.Query(getCommentsQuery)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotContent("comment is not exists")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	defer responses.Close()
	var comments []*entity.Comment
	for responses.Next() {
		var comment entity.Comment

		err = responses.Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				return nil, errs.NewNotContent("comment is not exists")
			}
			return nil, errs.NewInternalServerError("something went wrong")
		}

		comments = append(comments, &comment)
	}
	return comments, nil
}

func (commentPG *commentPG) GetCommentById(commentId int) (*entity.Comment, errs.MessageErr) {
	row := commentPG.db.QueryRow(getCommentByIdQuery, commentId)

	var comment entity.Comment

	err := row.Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotContent("comment is not exists")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &comment, nil
}

func (commentPG *commentPG) UpdateComment(payload entity.Comment) errs.MessageErr {
	_, err := commentPG.db.Exec(updateCommentQuery, payload.Message, payload.Id)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
func (commentPG *commentPG) DeleteComment(commentId int) errs.MessageErr {
	_, err := commentPG.db.Exec(deleteCommentQuery, commentId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
