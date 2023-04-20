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

func (h HttpServer) CreatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}
	photo.UserID = int(userID)

	res, err := h.app.CreatePhoto(photo)
	if err != nil {
		helpers.ErrorResponse(c, "create photo failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "create photo success", res, http.StatusOK)
}

func (h HttpServer) UpdatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	userID := uint(userData["id"].(float64))

	photo := models.Photo{}
	photo.UserID = int(userID)
	photo.Id = int(photoId)

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	res, err := h.app.UpdatePhoto(photo)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "update photo failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "update photo failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "update photo success", res, http.StatusOK)
}

func (h HttpServer) DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	err = h.app.DeletePhoto(int64(photoId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "delete photo failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "delete photo failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "delete photo success", nil, http.StatusOK)
}

func (h HttpServer) GetPhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.app.GetPhotoById(int64(photoId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "get photo failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "get photo failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "get photo success", res, http.StatusOK)
}

func (h HttpServer) GetAllPhoto(c *gin.Context) {

	res, err := h.app.GetAllPhoto()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "get all photo failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "get all photo failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "get all photo success", res, http.StatusOK)
}
