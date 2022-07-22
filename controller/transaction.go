package controller

import (
	"login-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransaction(c *gin.Context) {
	transaction, err := models.GetAllTransaction()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func EditTransaction(c *gin.Context) {
	transaction := models.Transaction{}
	err := c.ShouldBind(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = models.EditTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, http.StatusOK)
}
