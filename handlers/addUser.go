package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpopadic/go-fly/models"
)

func (h handler) AddUser(c *gin.Context) {
	var user models.User
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong params",
		})
	}

	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	c.JSON(http.StatusCreated, user)
}
