package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*DBO, error) {
	cfg := NewConfig()
	dbo, err := GetDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return dbo, nil
}

func NewConfig() *DBConfig {
	return &DBConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "hao",
		Password: "020899",
		DBName:   "gai-scraper",
	}
}

func (c *DBConfig) BuildConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

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

	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
