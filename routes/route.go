package routes

import (
	"login-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutesUser(route *gin.Engine) {

	aV1 := route.Group("/admin/v1")
	{
		aV1.GET("/get", controller.GetAllAdmin)
		aV1.POST("/getinfo", controller.GetInfoAdmin)
		aV1.POST("/add", controller.AddAdmin)
		aV1.POST("/edit", controller.EditAdmin)
		aV1.POST("/delete", controller.DeleteAdmin)
	}

	cV1 := route.Group("/category/v1")
	{
		cV1.GET("/get", controller.GetAllCategory)
		cV1.POST("/add", controller.AddCategory)
		cV1.POST("/edit", controller.EditCategory)
		cV1.POST("/delete", controller.DeleteCategory)
	}

	iV1 := route.Group("/image/v1")
	{
		iV1.POST("/add", controller.UploadProductImage)
	}

	oV1 := route.Group("/orders/v1")
	{
		oV1.GET("/get", controller.GetAllOrder)
		oV1.POST("/edit", controller.EditOrder)
	}

	uV1 := route.Group("/user/v1")
	{
		uV1.POST("/register", controller.UserRegister)
		uV1.POST("/login", controller.UserLogin)
		uV1.GET("/profile", controller.UserProfile)
	}

	pV1 := route.Group("/product/v1")
	{
		pV1.GET("/get", controller.GetAllProduct)
		pV1.POST("/getinfo", controller.GetInfoProduct)
		pV1.POST("/add", controller.AddProduct)
		pV1.POST("/edit", controller.EditProduct)
		pV1.POST("/delete", controller.DeleteProduct)
	}

	tV1 := route.Group("/transaction/v1")
	{
		tV1.GET("/get", controller.GetAllTransaction)
		tV1.POST("/edit", controller.EditTransaction)
	}
}
