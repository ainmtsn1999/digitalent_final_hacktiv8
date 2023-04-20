package handlers

import (
	"net/http"

	"github.com/ainmtsn1999/digitalent_final_hacktiv8/helpers"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	res, err := h.app.UserRegister(user)
	if err != nil {
		helpers.ErrorResponse(c, "register failed", err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SuccessResponse(c, "register success", res, http.StatusOK)
}

func (h HttpServer) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == helpers.AppJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	passRequest := user.Password

	res, err := h.app.UserLogin(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorResponse(c, "invalid email / username", err.Error(), http.StatusNotFound)
			return
		}
		helpers.ErrorResponse(c, "login failed", err.Error(), http.StatusBadRequest)
		return
	}

	isValid := helpers.ComparePass([]byte(res.Password), []byte(passRequest))
	if !isValid {
		helpers.ErrorResponse(c, "invalid password", helpers.ErrMap[http.StatusUnauthorized], http.StatusUnauthorized)
		return
	}

	token := helpers.GenerateToken(res.Id, res.Email, res.Username)

	helpers.SuccessResponse(c, "login success", gin.H{"token": token}, http.StatusAccepted)

}
