package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetConnect() *gorm.DB {
	dsn := "host=localhost user=admin password=root dbname=pgsql_db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return db
}
