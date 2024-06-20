package services

import (
	"github.com/labstack/echo"
	"github.com/otiai10/gosseract/v2"
	"gorm.io/gorm"
	"image-parser/app/models"
	"image-parser/app/repositories"
	"io"
	"mime/multipart"
	"os"
)

type ParseService struct {
	Db *gorm.DB
}

func (s *ParseService) Parse(c echo.Context) int {
	form, _ := c.MultipartForm()

	files := form.File["image"]

	for _, file := range files {
		go s.parseImage(file)
	}

	return 1
}

func (s *ParseService) parseImage(file *multipart.FileHeader) bool {
	gosseractClient := gosseract.NewClient()
	src, err := file.Open()
	if err != nil {
		return false
	}
	defer src.Close()

	newFile, err := os.Create("../storage/images/" + file.Filename)
	if err != nil {
		return false
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, src); err != nil {
		return false
	}
	var imagePath = "../storage/images/" + file.Filename
	gosseractClient.SetImage(imagePath)
	text, _ := gosseractClient.Text()

	image := models.Image{
		ImagePath:    imagePath,
		ParsedResult: text,
	}

	repositories.Сreate(s.Db, &image)

	return true
}
