package controller

import (
	"login-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAdmin(c *gin.Context) {
	admin, err := models.GetAllAdmin()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, admin)
}

func AddAdmin(c *gin.Context) {
	data := models.Admin{}
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if data.Name == "" {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}
	if data.UserName == "" {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}
	if data.PassWord == "" {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	err = models.AddAdmin(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}

func GetInfoAdmin(c *gin.Context) {
	binddata := models.Admin{}
	err := c.ShouldBind(&binddata)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	data, err := models.GetOneAdmin(binddata.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func EditAdmin(c *gin.Context) {
	data := models.Admin{}
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	if data.Name == "" {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	errs := models.EditAdmin(data)
	if errs != nil {
		c.JSON(http.StatusBadRequest, errs.Error())
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}

func DeleteAdmin(c *gin.Context) {
	data := models.Admin{}
	err := c.ShouldBind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	if data.ID == 0 {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	err = models.DeleteAdmin(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, http.StatusOK)
}
