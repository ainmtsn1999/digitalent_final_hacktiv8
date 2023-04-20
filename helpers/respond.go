package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	AppJSON = "application/json"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

var ErrMap = map[int]string{
	http.StatusBadRequest:          "bad request",
	http.StatusForbidden:           "forbidden access",
	http.StatusUnauthorized:        "unauthorized",
	http.StatusNotFound:            "not found",
	http.StatusUnprocessableEntity: "unprocessable entity",
	http.StatusInternalServerError: "internal server error",
}

func SuccessResponse(c *gin.Context, msg string, data interface{}, statusCode int) {
	c.JSON(statusCode, &Response{
		Message: msg,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, msg string, err string, statusCode int) {
	c.AbortWithStatusJSON(statusCode, &Response{
		Message: msg,
		Error:   err,
	})
}
