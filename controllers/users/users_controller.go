package users

import (
	"github.com/gin-gonic/gin"
	"github/aryabaguss/bookstore_users-api/domain/users"
	"github/aryabaguss/bookstore_users-api/services"
	"github/aryabaguss/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadrequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}