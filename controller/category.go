package controller

import (
	"login-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	category, err := models.GetAllCategory()
	if err != nil {
		return
	}
	c.JSON(200, category)
}
func AddCategory(c *gin.Context) {
	data := models.Category{}
	err := c.ShouldBind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if data.Name == "" {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	err = models.AddCategory(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}

func EditCategory(c *gin.Context) {
	data := models.Category{}
	err := c.ShouldBind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	if data.Name == "" {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	models.EditCategory(data)
	c.JSON(http.StatusOK, http.StatusOK)
}

func DeleteCategory(c *gin.Context) {
	data := models.Category{}
	err := c.ShouldBind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	if data.ID == 0 {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	models.DeleteCategory(data)
	c.JSON(http.StatusOK, http.StatusOK)
}
