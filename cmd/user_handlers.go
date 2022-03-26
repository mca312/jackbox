package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mca312/jackbox/server/app"
)

// RegisterHandler verifies the request and registers the user
func (m *Manager) RegisterHandler(c *gin.Context) {
	var registerUser app.User

	if err := c.BindJSON(&registerUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := m.Register(registerUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

// LoginHandler verifies the request and logs the user in
func (m *Manager) LoginHandler(c *gin.Context) {
	var loginUser app.User

	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := m.Login(loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

// UserHandler verifies the request and returns a user object
func (m *Manager) UserHandler(c *gin.Context) {
	ident := c.Param("user")
	if ident == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty user."})
		return
	}

	id, err := strconv.ParseInt(ident, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed id."})
		return
	}

	user, err := m.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id."})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
