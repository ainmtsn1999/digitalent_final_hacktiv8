package middlewares

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ainmtsn1999/digitalent_final_hacktiv8/config"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/helpers"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(photoId)

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := models.Photo{}

		err = db.First(&photo, uint(photoId)).Error
		log.Println(err)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				helpers.ErrorResponse(c, "photo auth not found", err.Error(), http.StatusNotFound)
				return
			}
			helpers.ErrorResponse(c, "photo auth bad request", err.Error(), http.StatusBadRequest)
			return
		}

		if photo.UserID != int(userID) {
			helpers.ErrorResponse(c, "you're not allowed to access this endpoint {PhotoAuth}", helpers.ErrMap[http.StatusUnauthorized], http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func CommentAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(commentId)

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = db.First(&comment, uint(commentId)).Error
		log.Println(err)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				helpers.ErrorResponse(c, "comment auth not found", err.Error(), http.StatusNotFound)
				return
			}
			helpers.ErrorResponse(c, "comment auth bad request", err.Error(), http.StatusBadRequest)
			return
		}

		if comment.UserID != int(userID) {
			helpers.ErrorResponse(c, "you're not allowed to access this endpoint {CommentAuth}", helpers.ErrMap[http.StatusUnauthorized], http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func SosmedAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		sosmedId, err := strconv.Atoi(c.Param("sosmedId"))
		if err != nil {
			helpers.ErrorResponse(c, "must be a valid id", err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(sosmedId)

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		sosmed := models.SocialMedia{}

		err = db.First(&sosmed, uint(sosmedId)).Error
		log.Println(err)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				helpers.ErrorResponse(c, "social media auth not found", err.Error(), http.StatusNotFound)
				return
			}
			helpers.ErrorResponse(c, "social media auth bad request", err.Error(), http.StatusBadRequest)
			return
		}

		if sosmed.UserID != int(userID) {
			helpers.ErrorResponse(c, "you're not allowed to access this endpoint {SosmedAuth}", helpers.ErrMap[http.StatusUnauthorized], http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
