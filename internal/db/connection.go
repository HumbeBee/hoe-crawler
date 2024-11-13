package db

import (
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
