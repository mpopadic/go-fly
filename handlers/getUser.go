package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpopadic/go-fly/models"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := h.DB.First(&user, 10).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"user_id": id, "status": "no value"})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
