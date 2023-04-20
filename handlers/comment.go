package handlers

import (
	"net/http"
	"strconv"

	"github.com/ainmtsn1999/digitalent_final_hacktiv8/helpers"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h HttpServer) CreateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	comment := models.Comment{}
	userID := userData["id"].(float64)

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}
	comment.UserID = int(userID)
	comment.PhotoID = photoId

	_, err = h.app.GetPhotoById(int64(photoId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "get photo for comment failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "get photo for comment failed", err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.app.CreateComment(comment)
	if err != nil {
		helpers.ErrorResponse(c, "create comment failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "create comment success", res, http.StatusOK)
}

func (h HttpServer) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	userID := userData["id"].(float64)

	comment := models.Comment{}
	comment.UserID = int(userID)
	comment.Id = int(commentId)

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	res, err := h.app.UpdateComment(comment)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "update comment failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "update comment failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "update comment success", res, http.StatusOK)
}

func (h HttpServer) DeleteComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	err = h.app.DeleteComment(int64(commentId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "delete comment failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "delete comment failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "delete comment success", nil, http.StatusOK)
}

func (h HttpServer) GetComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.app.GetCommentById(int64(commentId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "get comment failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "get comment failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "get comment success", res, http.StatusOK)
}

func (h HttpServer) GetAllComment(c *gin.Context) {

	res, err := h.app.GetAllComment()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "get all comment failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "get all comment failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "get all comment success", res, http.StatusOK)
}
