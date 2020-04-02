package users

import (
	"fmt"
	"github.com/AnthonyNixon/trivia-api/cmd/middleware"
	"github.com/AnthonyNixon/trivia-api/cmd/services/token"
	"github.com/AnthonyNixon/trivia-api/cmd/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUsersV1(router *gin.Engine) {
	router.Use(middleware.AttachInterfaceToGinContext(user.Service{}, user.ServiceKey))
	router.Use(middleware.AttachInterfaceToGinContext(token.Service{}, token.ServiceKey))
	router.POST("/v1/users", NewUser)
}

func NewUser(c *gin.Context) {
	var newUser user.User

	bindErr := c.BindJSON(&newUser)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Bad JSON Input: %s", bindErr.Error()))
		return
	}

	userService := middleware.GetInterfaceFromGinContext(c, user.ServiceKey).(user.Interface)

	createdUser, err := userService.CreateNewUser(newUser.Name)
	if err != nil {
		c.JSON(err.StatusCode(), err.Description())
		return
	}

	tokenService := middleware.GetInterfaceFromGinContext(c, token.ServiceKey).(token.Interface)
	token, error := tokenService.NewUserToken(createdUser)
	if error != nil {
		c.JSON(error.StatusCode(), error.Description())
	}

	c.Header("user-token", token)
	c.JSON(http.StatusCreated, createdUser)
}
