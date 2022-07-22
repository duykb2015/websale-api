package controller

import (

	// "upfile/config"
	// "upfile/utils/filecheck"
	// "upfile/utils/respo"

	"image/jpeg"
	"io"
	"log"
	"login-api/models"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func UploadProductImage(c *gin.Context) {

	c.Request.ParseMultipartForm(30 << 40)

	form, _ := c.MultipartForm()
	formfile := form.File["file"]

	if formfile == nil {
		c.JSON(400, "File not found")
		return
	}

	imagesize := models.ImageSize{}
	c.Bind(&imagesize)

	file, err := formfile[0].Open()

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	filename := "img-" + strconv.Itoa(randNum()) + ".png"

	if imagesize.Height == 0 || imagesize.Width == 0 {

		defer file.Close()
		out, err := os.Create("../websales/upload/product/" + filename)
		if err != nil {
			c.JSON(400, err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)

		if err != nil {
			c.JSON(400, err.Error())
			return
		}
	} else {
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		m := resize.Resize(uint(imagesize.Width), uint(imagesize.Height), img, resize.Lanczos3)

		out, err := os.Create("../websales/upload/product/" + filename)

		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		jpeg.Encode(out, m, nil)
	}

	c.JSON(http.StatusOK, filename)
}

func randNum() int {
	var max = 10000000
	var min = 10000
	return rand.Intn(max) + rand.Intn(min)
}
