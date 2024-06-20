package repositories

import (
	"gorm.io/gorm"
	"image-parser/app/models"
	"time"
)

func Сreate(db *gorm.DB, model *models.Image) {
	model.CreatedAt = time.Now()
	db.Create(model)
}
