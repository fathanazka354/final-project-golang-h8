package handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/helpers"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) commentHandler {
	return commentHandler{
		commentService: commentService,
	}
}

func (ch *commentHandler) CreateComment(ctx *gin.Context) {
	var nerCommentRequest dto.NewCommentRequest

	if err := ctx.ShouldBindJSON(&nerCommentRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := ch.commentService.CreateComment(nerCommentRequest)

	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (ch *commentHandler) GetComments(ctx *gin.Context) {

	comments, err := ch.commentService.GetComments()

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(comments.StatusCode, comments)
}

func (ch *commentHandler) GetCommentById(ctx *gin.Context) {
	commentId, err := helpers.GetParamById(ctx, "commentId")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	comment, err := ch.commentService.GetCommentById(commentId)

	if err != nil {
		ctx.JSON(comment.StatusCode, comment)
		return
	}

	ctx.JSON(comment.StatusCode, comment)
}
func (ch *commentHandler) UpdateComment(ctx *gin.Context) {
	commentId, err := helpers.GetParamById(ctx, "commentId")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	var newCommentRequest dto.NewCommentRequest
	if err := ctx.ShouldBindJSON(&newCommentRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid request body",
		})
		return
	}
	comment, err := ch.commentService.UpdateComment(commentId, user.Id, newCommentRequest.PhotoId, newCommentRequest)

	if err != nil {
		ctx.JSON(comment.StatusCode, comment)
		return
	}

	ctx.JSON(comment.StatusCode, comment)
}
func (ch *commentHandler) DeleteComment(ctx *gin.Context) {
	commentId, err := helpers.GetParamById(ctx, "commentId")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	comment, err := ch.commentService.DeleteComment(commentId)

	if err != nil {
		ctx.JSON(comment.StatusCode, comment)
		return
	}

	ctx.JSON(comment.StatusCode, comment)
}
