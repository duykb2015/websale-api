package controller

import (
	"login-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	if product, err := models.GetAllProduct(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetInfoProduct(c *gin.Context) {
	binddata := models.Product{}
	err := c.ShouldBind(&binddata)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	data, err := models.GetOneProduct(binddata.Slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func AddProduct(c *gin.Context) {
	data := models.Product{}
	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if data.Name == "" {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	err = models.AddProduct(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}

func EditProduct(c *gin.Context) {
	data := models.Product{}
	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = models.EditProduct(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}

func DeleteProduct(c *gin.Context) {
	data := models.Product{}
	err := c.ShouldBind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	if data.ID == 0 {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	models.DeleteProduct(data)
	c.JSON(http.StatusOK, http.StatusOK)
}
