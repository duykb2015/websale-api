package controller

import (
	"login-api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Xử lí đăng nhập cho người dùng
func UserRegister(c *gin.Context) {
	u := models.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if !models.IsEmailValid(u.Email) {
		c.JSON(http.StatusBadRequest, "Please enter an email")
		return
	}

	if _, isExist := models.CheckEmailExist(u.Email); isExist {
		c.JSON(http.StatusBadRequest, "This email had already exist, please try another email")
		return
	}

	if u.Email == "" || u.Password == "" || u.PasswordConfirm == "" {
		c.JSON(http.StatusBadRequest, "Please fill out the information completely")
		return
	}

	if u.Password != u.PasswordConfirm {
		c.JSON(http.StatusBadRequest, "Password and password confirm do not match")
		return
	}

	models.CreateNewUser(u)
}

func UserLogin(c *gin.Context) {
	u := models.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if ut, isExist := models.CheckEmailExist(u.Email); !isExist {
		c.JSON(http.StatusBadRequest, "This email does not exist, please check again")
		return
	} else {
		if u.Password != ut.Password {
			c.JSON(http.StatusBadRequest, "Wrong password, please check again")
			return
		}
		token, err := models.CreateToken(ut.ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, token)
	}
}

func UserProfile(c *gin.Context) {
	bearToken := c.GetHeader("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 0 {
		panic("error")
	}

	if u, err := models.GetInfoUser(c.Request, strArr[1]); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, u)
	}
}
