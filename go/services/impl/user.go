package impl

import (
	// Go native packages
	"net/http"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"

	// Go native packages
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Password string `json:"password" form:"password" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		env.Environment.Logger.Error("Bind json error" + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := env.Environment.DataStore.SelectUserWithUsername(c, request.Username)
	if err != nil {
		params := make(map[string]interface{})
		params["username"] = request.Username
		msg := "Get user password error: " + err.Error()
		env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(msg)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if user.Password != request.Password {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := env.Environment.Middlewares.GenerateToken(user.Id)
	if err != nil {
		params := make(map[string]interface{})
		params["user-id"] = user.Id
		msg := "Genereate token error: " + err.Error()
		env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(msg)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
