package middlewares

import (
	"net/http"

	"github.com/ainmtsn1999/digitalent_final_hacktiv8/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			helpers.ErrorResponse(c, helpers.ErrMap[http.StatusForbidden], err.Error(), http.StatusForbidden)
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
