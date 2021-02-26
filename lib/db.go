package lib

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	Message   string
	UserAgent string
}

func DbConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func DbInit() {
	db := DbConn()
	db.AutoMigrate(&Tweet{})
}
