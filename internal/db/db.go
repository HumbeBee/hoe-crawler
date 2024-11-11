package db

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB(cfg *DBConfig) (*DBO, error) {
	db, err := gorm.Open(mysql.Open(cfg.BuildConnectionString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBO{db: db}, nil
}

func (dbo *DBO) Close() {
	sqlDB, err := dbo.db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func (dbo *DBO) InsertHoe(hoeInfo *models.HoeInfo) error {
	transaction := dbo.db.Begin()

	if err := transaction.Create(hoeInfo).Error; err != nil {
		transaction.Rollback()
		return err
	}

	if err := transaction.Commit().Error; err != nil {
		return err
	}

	fmt.Println("Inserted")
	hoeInfo.Print()

	return nil
}
