package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mca312/jackbox/server/internal"
)

type UserModel struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

func RegisterHandler(c *gin.Context) {
	var registerUser UserModel

	if err := c.BindJSON(&registerUser); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("/register - bad reqeust: %+v", err))
		return
	}

	// check email exists and has one @ symbol
	if !internal.IsValidEmail(registerUser.Email) {
		c.JSON(http.StatusBadRequest, "/register - invalid email")
		return
	}

	if err := internal.Register(internal.User(registerUser)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}

func LoginHandler(c *gin.Context) {
	var loginUser UserModel

	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("/login - bad reqeust: %+v", err))
		return
	}

	user, err := internal.Login(internal.User(loginUser))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func UserHandler(c *gin.Context) {
	ident := c.Param("user")
	if ident == "" {
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	id, err := strconv.ParseInt(ident, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request, id must be an int")
		return
	}

	user, err := internal.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Invalid request for id: %d", id))
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
