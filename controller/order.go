package controller

import (
	"login-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrder(c *gin.Context) {
	order, err := models.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, order)
}

func EditOrder(c *gin.Context) {
	order := models.Order{}
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = models.EditOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}
