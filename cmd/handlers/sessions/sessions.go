package sessions

import (
	"fmt"
	"net/http"

	"github.com/AnthonyNixon/trivia-api/cmd/middleware"
	"github.com/AnthonyNixon/trivia-api/cmd/services/session"
	"github.com/gin-gonic/gin"
)

func AddSessionsV1(router *gin.Engine) {
	router.Use(middleware.AttachInterfaceToGinContext(session.Service{}, session.ServiceKey))
	router.POST("/v1/sessions", NewSession)
}

func NewSession(c *gin.Context) {
	var newSession session.Session

	bindErr := c.BindJSON(&newSession)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Bad JSON Input: %s", bindErr.Error()))
		return
	}

	sessionService := middleware.GetInterfaceFromGinContext(c, session.ServiceKey).(session.Interface)

	createdSession, err := sessionService.CreateNewSession(newSession.Name, newSession.Description, newSession.StartTime)
	if err != nil {
		c.JSON(err.StatusCode(), err.Description())
		return
	}

	c.JSON(http.StatusCreated, createdSession)
}
