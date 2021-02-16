package controllers

import (
	"net/http"
	"time"

	m "github.com/Flix14/casbin-example-2/models"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Name string `json:"name"`
}

func Signin(c *gin.Context) {
	var creds Credentials
	user := new(m.User)

	err := c.BindJSON(&creds)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	user, err = user.VerifyCredentials(creds.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}

	token, exp, err := user.GenerateTokenJWT()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: exp,
	})
	c.JSON(200, gin.H{"user": user, "token": token})
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}
