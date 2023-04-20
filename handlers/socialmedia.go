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

func (h HttpServer) CreateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	socialmedia := models.SocialMedia{}
	userID := userData["id"].(float64)

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&socialmedia)
	} else {
		c.ShouldBind(&socialmedia)
	}
	socialmedia.UserID = int(userID)

	res, err := h.app.CreateSocialMedia(socialmedia)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			helpers.ErrorResponse(c, "user has been made a social media", err.Error(), http.StatusUnprocessableEntity)
			return

		}
		helpers.ErrorResponse(c, "create social media failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "create social media success", res, http.StatusOK)
}

func (h HttpServer) UpdateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	sosmedId, err := strconv.Atoi(c.Param("sosmedId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	userID := userData["id"].(float64)

	socialmedia := models.SocialMedia{}
	socialmedia.UserID = int(userID)
	socialmedia.Id = int(sosmedId)

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&socialmedia)
	} else {
		c.ShouldBind(&socialmedia)
	}

	res, err := h.app.UpdateSocialMedia(socialmedia)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "update social media failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "update social media failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "update social media success", res, http.StatusOK)
}

func (h HttpServer) DeleteSocialMedia(c *gin.Context) {
	sosmedId, err := strconv.Atoi(c.Param("sosmedId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	err = h.app.DeleteSocialMedia(int64(sosmedId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "delete social media failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "delete social media failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "delete social media success", nil, http.StatusOK)
}

func (h HttpServer) GetSocialMedia(c *gin.Context) {
	sosmedId, err := strconv.Atoi(c.Param("sosmedId"))
	if err != nil {
		helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.app.GetSocialMediaById(int64(sosmedId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "get social media failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "get social media failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "get social media success", res, http.StatusOK)
}

func (h HttpServer) GetAllSocialMedia(c *gin.Context) {

	res, err := h.app.GetAllSocialMedia()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "get all social media failed", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "get all social media failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "get all social media success", res, http.StatusOK)
}
