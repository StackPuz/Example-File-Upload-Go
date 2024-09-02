package main

import (
	"app/models"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	router := gin.Default()
	uploadPath := "./public/uploads"
	os.MkdirAll(uploadPath, os.ModePerm)
	router.Static("/uploads", uploadPath)
	router.StaticFile("/", "./public/index.html")
	router.POST("/submit", func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindWith(&product, binding.FormMultipart); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		image, _ := c.FormFile("Image")
		filePath := filepath.Join(uploadPath, image.Filename)
		src, _ := image.Open()
		dst, _ := os.Create(filePath)
		io.Copy(dst, src)
		c.JSON(http.StatusOK, gin.H{"Name": product.Name, "Image": image.Filename})
	})
	router.Run()
}
